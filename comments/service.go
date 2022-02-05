package comments

import (
	"context"
	"log"
)

type ServiceComment struct {
	repo ICommentsRepository
}

func NewServiceComment(repo ICommentsRepository) *ServiceComment {
	return &ServiceComment{repo}
}

// func (s *ServiceComment) GetData(ctx context.Context) (data CtrData, err error) {
// 	getData, err := s.repo.GetData(ctx)
// 	if err != nil {
// 		log.Println(err)
// 		return data, err
// 	}

// 	return getData, nil

// }

func (s *ServiceComment) GetCurrentUser(ctx context.Context) (data CtrCurrentUser, err error) {
	getCurrentUser, err := s.repo.GetCurrentUser(ctx)
	if err != nil {
		log.Println(err)
		return data, err
	}

	return getCurrentUser, nil
}

func (s *ServiceComment) GetComments(ctx context.Context) (data []CtrComment, err error) {
	commentsData, err := s.repo.GetComments(ctx)
	if err != nil {
		log.Println(err)
		return data, err
	}

	repliesData, err := s.repo.GetReplies(ctx)
	if err != nil {
		log.Println(err)
		return data, err
	}

	users, err := s.repo.GetUser(ctx)
	if err != nil {
		log.Println(err)
		return data, err
	}

	currentUser, err := s.repo.GetCurrentUser(ctx)
	if err != nil {
		log.Println(err)
		return data, err
	}

	for _, comment := range commentsData {
		var dataComment CtrComment
		dataComment.ID = comment.ID
		dataComment.Content = comment.Content
		dataComment.CreatedAt = comment.CreatedAt
		dataComment.Score = comment.Score
		dataComment.Username = comment.Username
		dataComment.Image = comment.Image
		// var users []string

		for _, reply := range repliesData {
			var dataReply CtrReply
			dataReply.ID = reply.ID
			dataReply.Content = reply.Content
			dataReply.CreatedAt = reply.CreatedAt
			dataReply.Score = reply.Score
			dataReply.ReplyingTo = reply.ReplyingTo
			dataReply.Username = reply.Username
			dataReply.IdComment = reply.IdComment

			for _, user := range users {
				if user.Username == reply.Username {
					dataReply.Image = user.Image
				}
			}

			if currentUser.Username == reply.Username {
				dataReply.Image = currentUser.Image
			}

			if comment.ID == reply.IdComment {
				dataComment.Replies = append(dataComment.Replies, dataReply)
			}

			// if comment.Username == reply.ReplyingTo {
			// 	dataComment.Replies = append(dataComment.Replies, dataReply)
			// 	// users = append(users, reply.Username)
			// }

			// for _, reply2 := range dataComment.Replies {
			// 	if reply2.Username == reply.ReplyingTo {
			// 		dataComment.Replies = append(dataComment.Replies, dataReply)
			// 	}
			// }

		}

		data = append(data, dataComment)
	}

	return data, nil

}

func (s *ServiceComment) GetReplies(ctx context.Context) (data []Reply, err error) {
	repliesData, err := s.repo.GetReplies(ctx)
	if err != nil {
		log.Println(err)
		return data, err
	}

	return repliesData, nil

}

func (s *ServiceComment) PostComment(ctx context.Context, data CtrPostComment) error {
	err := s.repo.PostComment(ctx, data)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func (s *ServiceComment) PostReply(ctx context.Context, data CtrPostReply) error {
	err := s.repo.PostReply(ctx, data)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

// func (s *ServiceComment) EditComment(ctx context.Context, id uint64, data Comment) error {
// 	err := s.repo.EditComment(ctx, id, data)
// 	if err != nil {
// 		log.Println(err)
// 		return err
// 	}

// 	return nil
// }

// func (s *ServiceComment) EditReply(ctx context.Context, id uint64, data Reply) error {
// 	err := s.repo.EditReply(ctx, id, data)
// 	if err != nil {
// 		log.Println(err)
// 		return err
// 	}

// 	return nil
// }

func (s *ServiceComment) DeleteComment(ctx context.Context, id uint64) error {
	err := s.repo.DeleteComment(ctx, id)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func (s *ServiceComment) DeleteReply(ctx context.Context, id uint64) error {
	err := s.repo.DeleteReply(ctx, id)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func (s *ServiceComment) EditCommentContent(ctx context.Context, id uint64, data CtrEditComment) error {
	err := s.repo.EditCommentContent(ctx, id, data)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func (s *ServiceComment) EditReplyContent(ctx context.Context, id uint64, data CtrEditReply) error {
	err := s.repo.EditReplyContent(ctx, id, data)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func (s *ServiceComment) EditReplyScore(ctx context.Context, id uint64, data CtrEditScore) error {
	err := s.repo.EditReplyScore(ctx, id, data)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func (s *ServiceComment) EditCommentScore(ctx context.Context, id uint64, data CtrEditScore) error {
	err := s.repo.EditCommentScore(ctx, id, data)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}
