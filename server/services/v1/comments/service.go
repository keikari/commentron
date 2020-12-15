package comments

import (
	"database/sql"
	"math"
	"net/http"

	"github.com/lbryio/commentron/commentapi"
	m "github.com/lbryio/commentron/model"
	"github.com/lbryio/commentron/server/lbry"

	"github.com/lbryio/lbry.go/extras/api"
	"github.com/lbryio/lbry.go/v2/extras/errors"
	"github.com/lbryio/lbry.go/v2/extras/util"
	v "github.com/lbryio/ozzo-validation"

	"github.com/volatiletech/null"
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries/qm"
)

/*
'ping': ping,
DONE 'get_claim_comments': handle_get_claim_comments,  # this gets used ( params: claimid, page, page-size, include-replies, order-by )
DONE 'get_channel_from_comment_id': handle_get_channel_from_comment_id,  # this gets used by SDK
DONE 'create_comment': handle_create_comment,  # this gets used
'abandon_comment': handle_abandon_comment,  # this gets used
'edit_comment': handle_edit_comment  # this gets used
NEW APIS
comment count per claim
comment count per parent
comments per parent ( order-by time, rating, page, page-size )
page for comment ( params: page, [size], order-by )

*/

// Service is the service struct defined for the comment package for rpc service "comment.*"
type Service struct{}

// Create creates a comment
func (c *Service) Create(_ *http.Request, args *commentapi.CreateArgs, reply *commentapi.CreateResponse) error {
	err := v.ValidateStruct(args,
		v.Field(&args.ClaimID, v.Required))
	if err != nil {
		return api.StatusError{Err: errors.Err(err), Status: http.StatusBadRequest}
	}
	channel, err := m.Channels(m.ChannelWhere.ClaimID.EQ(null.StringFromPtr(args.ChannelID).String)).OneG()
	if errors.Is(err, sql.ErrNoRows) {
		channel = &m.Channel{
			ClaimID: null.StringFromPtr(args.ChannelID).String,
			Name:    null.StringFromPtr(args.ChannelName).String,
		}
		err = nil
		err := channel.InsertG(boil.Infer())
		if err != nil {
			return errors.Err(err)
		}
	}
	blockedEntry, err := m.BlockedEntries(m.BlockedEntryWhere.UniversallyBlocked.EQ(null.BoolFrom(true)), m.BlockedEntryWhere.BlockedChannelID.EQ(null.StringFromPtr(args.ChannelID))).OneG()
	if err != nil {
		return errors.Err(err)
	}
	if blockedEntry != nil {
		return api.StatusError{Err: errors.Err("channel is not allowed to post comments"), Status: http.StatusBadRequest}
	}
	if err != nil {
		return errors.Err(err)
	}
	commentID, timestamp, err := createCommentID(args.CommentText, null.StringFromPtr(args.ChannelID).String)
	if err != nil {
		return errors.Err(err)
	}

	comment, err := m.Comments(m.CommentWhere.CommentID.EQ(commentID)).OneG()
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return errors.Err(err)
	}

	if comment != nil {
		return api.StatusError{Err: errors.Err("duplicate comment!"), Status: http.StatusBadRequest}
	}
	signingChannel, err := lbry.GetSigningChannelForClaim(args.ClaimID)
	if err != nil {
		return errors.Err(err)
	}
	if signingChannel != nil {
		blockedEntry, err := m.BlockedEntries(m.BlockedEntryWhere.BlockedByChannelID.EQ(null.StringFrom(signingChannel.ClaimID)), m.BlockedEntryWhere.BlockedChannelID.EQ(null.StringFromPtr(args.ChannelID))).OneG()
		if err != nil {
			return errors.Err(err)
		}
		if blockedEntry != nil {
			return api.StatusError{Err: errors.Err("channel %s is blocked by publisher %s", args.ChannelID, signingChannel.Name)}
		}
	}

	comment = &m.Comment{
		CommentID:   commentID,
		LbryClaimID: args.ClaimID,
		ChannelID:   null.StringFromPtr(args.ChannelID),
		Body:        args.CommentText,
		ParentID:    null.StringFromPtr(args.ParentID),
		Signature:   null.StringFromPtr(args.Signature),
		Signingts:   null.StringFromPtr(args.SigningTS),
		Timestamp:   int(timestamp),
	}

	err = errors.Err(comment.InsertG(boil.Infer()))
	if err != nil {
		return errors.Err(err)
	}
	item := populateItem(comment, channel)
	reply.CommentItem = &item

	go lbry.Notify(lbry.NotifyOptions{
		ActionType: "C",
		CommentID:  item.CommentID,
		ChannelID:  &item.ChannelID,
		ParentID:   &item.ParentID,
		Comment:    &item.Comment,
		ClaimID:    item.ClaimID,
	})
	return nil
}

// List lists comments based on filters and arguments passed. The returned result is dynamic based on the args passed
func (c *Service) List(_ *http.Request, args *commentapi.ListArgs, reply *commentapi.ListResponse) error {
	args.ApplyDefaults()
	loadChannels := qm.Load("Channel")
	filterIsHidden := m.CommentWhere.IsHidden.EQ(null.BoolFrom(true))
	filterClaimID := m.CommentWhere.LbryClaimID.EQ(util.StrFromPtr(args.ClaimID))
	filterAuthorClaimID := m.CommentWhere.ChannelID.EQ(null.StringFromPtr(args.AuthorClaimID))
	filterTopLevel := m.CommentWhere.ParentID.IsNull()
	filterParent := m.CommentWhere.ParentID.EQ(null.StringFrom(util.StrFromPtr(args.ParentID)))

	totalCommentsQuery := make([]qm.QueryMod, 0)
	offset := (args.Page - 1) * args.PageSize
	getCommentsQuery := []qm.QueryMod{loadChannels, qm.Offset(offset), qm.Limit(args.PageSize), qm.OrderBy(m.CommentColumns.Timestamp + " DESC")}
	hasHiddenCommentsQuery := []qm.QueryMod{filterIsHidden, qm.Limit(1)}

	if args.AuthorClaimID != nil {
		getCommentsQuery = append(getCommentsQuery, filterAuthorClaimID)
		hasHiddenCommentsQuery = append(hasHiddenCommentsQuery, filterAuthorClaimID)
		totalCommentsQuery = append(totalCommentsQuery, filterAuthorClaimID)
	}

	if args.ClaimID != nil {
		getCommentsQuery = append(getCommentsQuery, filterClaimID)
		hasHiddenCommentsQuery = append(hasHiddenCommentsQuery, filterClaimID)
		totalCommentsQuery = append(totalCommentsQuery, filterClaimID)
	}

	if args.TopLevel {
		getCommentsQuery = append(getCommentsQuery, filterTopLevel)
		hasHiddenCommentsQuery = append(hasHiddenCommentsQuery, filterTopLevel)
		totalCommentsQuery = append(totalCommentsQuery, filterTopLevel)
	}

	if args.ParentID != nil {
		getCommentsQuery = append(getCommentsQuery, filterParent)
		hasHiddenCommentsQuery = append(hasHiddenCommentsQuery, filterParent)
		totalCommentsQuery = append(totalCommentsQuery, filterParent)
	}

	totalItems, err := m.Comments(totalCommentsQuery...).CountG()
	if err != nil {
		return errors.Err(err)
	}

	hasHiddenComments, err := m.Comments(hasHiddenCommentsQuery...).ExistsG()
	if err != nil {
		return errors.Err(err)
	}

	comments, err := m.Comments(getCommentsQuery...).AllG()
	if err != nil {
		return errors.Err(err)
	}

	var items []commentapi.CommentItem
	for _, comment := range comments {
		var channel *m.Channel
		if comment.R != nil {
			channel = comment.R.Channel
			if channel != nil && channel.Name != "" {
				items = append(items, populateItem(comment, channel))
			}
		}
	}

	reply.Items = items
	reply.Page = args.Page
	reply.PageSize = args.PageSize
	reply.TotalItems = totalItems
	reply.TotalPages = int(math.Ceil(float64(totalItems) / float64(args.PageSize)))
	reply.HasHiddenComments = hasHiddenComments

	return nil
}

// GetChannelFromCommentID gets the channel info for a specific comment, this is really only used by the sdk
func (c *Service) GetChannelFromCommentID(_ *http.Request, args *commentapi.ChannelArgs, reply *commentapi.ChannelResponse) error {
	comment, err := m.Comments(m.CommentWhere.CommentID.EQ(args.CommentID), qm.Load(m.CommentRels.Channel)).OneG()
	if errors.Is(err, sql.ErrNoRows) {
		return api.StatusError{Err: errors.Err("could not find comment for comment id"), Status: http.StatusBadRequest}
	}
	if err != nil {
		return errors.Err(err)
	}
	if comment.R == nil && comment.R.Channel == nil {
		return api.StatusError{Err: errors.Err("could not find channel for comment"), Status: http.StatusBadRequest}
	}
	reply.ChannelID = comment.R.Channel.ClaimID
	reply.ChannelName = comment.R.Channel.Name
	return nil
}

// Abandon deletes a comment
func (c *Service) Abandon(_ *http.Request, args *commentapi.AbandonArgs, reply *commentapi.AbandonResponse) error {
	item, err := abandon(args)
	if err != nil {
		return errors.Err(err)
	}
	reply.CommentItem = item
	reply.Abandoned = true

	go lbry.Notify(lbry.NotifyOptions{
		ActionType: "D",
		CommentID:  item.CommentID,
		ChannelID:  &item.ChannelID,
		ParentID:   &item.ParentID,
		Comment:    &item.Comment,
		ClaimID:    item.ClaimID,
	})

	return nil
}

// Edit edits a comment
func (c *Service) Edit(_ *http.Request, args *commentapi.EditArgs, reply *commentapi.EditResponse) error {
	item, err := edit(args)
	if err != nil {
		return errors.Err(err)
	}
	reply.CommentItem = item

	go lbry.Notify(lbry.NotifyOptions{
		ActionType: "U",
		CommentID:  item.CommentID,
		ChannelID:  &item.ChannelID,
		ParentID:   &item.ParentID,
		Comment:    &item.Comment,
		ClaimID:    item.ClaimID,
	})
	return nil
}

// ByID returns the comment from the comment id passed in
func (c *Service) ByID(r *http.Request, args *commentapi.ByIDArgs, reply *commentapi.ByIDResponse) error {
	item, err := byID(r, args)
	if err != nil {
		return err
	}
	reply.Item = item
	return nil
}

// Pin sets the pinned flag on a comment
func (c *Service) Pin(r *http.Request, args *commentapi.PinArgs, reply *commentapi.PinResponse) error {
	item, err := pin(r, args)
	if err != nil {
		return err
	}
	reply.Item = item
	return nil
}
