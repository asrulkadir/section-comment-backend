package comments

import (
	"context"
	"database/sql"
	"log"
	"time"
)

type ICommentsRepository interface {
	// GetData(ctx context.Context) (data CtrData, err error)
	GetCurrentUser(ctx context.Context) (data CtrCurrentUser, err error)
	GetComments(ctx context.Context) (data []Comment, err error)
	GetReplies(ctx context.Context) (data []Reply, err error)
	GetUser(ctx context.Context) (data []CurrentUser, err error)
	PostComment(ctx context.Context, data CtrPostComment) error
	PostReply(ctx context.Context, data CtrPostReply) error
	// 	EditComment(ctx context.Context, id uint64, data Comment) error
	// 	EditReply(ctx context.Context, id uint64, data Reply) error
	DeleteComment(ctx context.Context, id uint64) error
	DeleteReply(ctx context.Context, id uint64) error
	EditCommentContent(ctx context.Context, id uint64, data CtrEditComment) error
	EditReplyContent(ctx context.Context, id uint64, data CtrEditReply) error
	EditReplyScore(ctx context.Context, id uint64, data CtrEditScore) error
	EditCommentScore(ctx context.Context, id uint64, data CtrEditScore) error
}

type commentRepository struct {
	db *sql.DB
}

func NewRepositoryComments(db *sql.DB) *commentRepository {
	return &commentRepository{db}
}

// func (c *commentRepository) GetData(ctx context.Context) (data CtrData, err error) {
// 	var currentUser CurrentUser
// 	err = c.db.QueryRowContext(ctx, "SELECT id, image, username FROM tbl_current_user WHERE id = 1").Scan(
// 		&currentUser.ID,
// 		&currentUser.Image,
// 		&currentUser.Username,
// 	)
// 	if err != nil {
// 		log.Println(err)
// 		return data, err
// 	}
// 	data.CurrentUser = CtrCurrentUser(currentUser)

// 	var comments []Comment
// 	queryComment := `SELECT id, content, created_at, score, username FROM tbl_comments`
// 	rows, err := c.db.QueryContext(ctx, queryComment)
// 	if err != nil {
// 		log.Println(err)
// 		return data, err
// 	}

// 	defer rows.Close()

// 	for rows.Next() {
// 		var comment Comment
// 		err := rows.Scan(&comment.ID, &comment.Content, &comment.CreatedAt, &comment.Score, &comment.Username)
// 		if err != nil {
// 			log.Println(err)
// 			return data, err
// 		}
// 		comments = append(comments, comment)
// 	}
// 	data.Comments = comments

// 	return data, nil

// }

// func (c *commentRepository) GetUsers(ctx context.Context) (data []CurrentUser, err error) {
// 	rows, err := c.db.QueryContext(ctx, "SELECT id, image, username FROM tbl_user")
// 	if err != nil {
// 		log.Println(err)
// 		return data, err
// 	}

// 	for rows.Next() {
// 		var user CurrentUser
// 		err = rows.Scan(&user.ID, &user.Image, &user.Username)
// 		if err != nil {
// 			log.Println(err)
// 			return data, err
// 		}

// 		data = append(data, user)
// 	}

// 	return data, nil
// }

// func (c *commentRepository) GetUsersByID(ctx context.Context, id uint64) (data CurrentUser, err error) {
// 	err = c.db.QueryRowContext(ctx, "SELECT id, image, username FROM tbl_user WHERE id = ?", id).Scan(&data.ID, &data.Image, &data.Username)
// 	if err != nil {
// 		log.Println(err)
// 		return data, err
// 	}

// 	return data, nil
// }

func (c *commentRepository) GetCurrentUser(ctx context.Context) (data CtrCurrentUser, err error) {
	err = c.db.QueryRowContext(ctx, "SELECT id, image, username FROM tbl_current_user WHERE id = 1").Scan(&data.ID, &data.Image, &data.Username)
	if err != nil {
		log.Println(err)
		return data, err
	}

	return data, nil
}

func (c *commentRepository) GetUser(ctx context.Context) (data []CurrentUser, err error) {
	query := `SELECT id, image, username FROM tbl_user`
	rows, err := c.db.QueryContext(ctx, query)
	if err != nil {
		log.Println(err)
		return data, err
	}

	defer rows.Close()

	for rows.Next() {
		var user CurrentUser
		err := rows.Scan(&user.ID, &user.Image, &user.Username)
		if err != nil {
			log.Println(err)
			return data, err
		}

		data = append(data, user)
	}

	return data, nil
}

func (c *commentRepository) GetComments(ctx context.Context) (data []Comment, err error) {
	query := `SELECT tc.id, tc.content, tc.created_at, tc.score, tc.username, tu.image  
			FROM tbl_comments tc
			JOIN tbl_user tu 
			ON tc.username = tu.username
			ORDER BY created_at ASC`
	rows, err := c.db.QueryContext(ctx, query)
	if err != nil {
		log.Println(err)
		return data, err
	}

	defer rows.Close()

	for rows.Next() {
		var comment Comment
		err := rows.Scan(&comment.ID,
			&comment.Content,
			&comment.CreatedAt,
			&comment.Score,
			&comment.Username,
			&comment.Image,
		)
		if err != nil {
			log.Println(err)
			return data, err
		}
		data = append(data, comment)
	}

	return data, nil
}

func (c *commentRepository) GetReplies(ctx context.Context) (data []Reply, err error) {
	query := `SELECT tr.id, tr.content, tr.created_at, tr.score, tr.replying_to, tr.username  
		FROM tbl_replies tr 
		ORDER BY created_at ASC`
	rows, err := c.db.QueryContext(ctx, query)
	if err != nil {
		log.Println(err)
		return data, err
	}
	defer rows.Close()
	for rows.Next() {
		var reply Reply
		err := rows.Scan(&reply.ID,
			&reply.Content,
			&reply.CreatedAt,
			&reply.Score,
			&reply.ReplyingTo,
			&reply.Username,
		)
		if err != nil {
			log.Println(err)
			return data, err
		}
		data = append(data, reply)
	}

	return data, nil
}

func (c *commentRepository) PostReply(ctx context.Context, data CtrPostReply) error {
	query := `INSERT INTO tbl_replies (content, score, replying_to, username) VALUES ($1, $2, $3, $4)`
	_, err := c.db.ExecContext(ctx, query, data.Content, data.Score, data.ReplyingTo, data.Username)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func (c *commentRepository) PostComment(ctx context.Context, data CtrPostComment) error {
	query := `INSERT INTO tbl_comments (content, score, username, created_at) VALUES ($1, $2, $3, $4)`
	_, err := c.db.ExecContext(ctx, query, data.Content, data.Score, data.Username, time.Now())
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func (c *commentRepository) DeleteComment(ctx context.Context, id uint64) error {
	query := `DELETE FROM tbl_comments WHERE id = $1`
	_, err := c.db.ExecContext(ctx, query, id)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func (c *commentRepository) DeleteReply(ctx context.Context, id uint64) error {
	query := `DELETE FROM tbl_replies WHERE id = $1`
	_, err := c.db.ExecContext(ctx, query, id)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func (c *commentRepository) EditCommentContent(ctx context.Context, id uint64, data CtrEditComment) error {
	query := `UPDATE tbl_comments SET content = $1 WHERE id = $2`
	_, err := c.db.ExecContext(ctx, query, data.Content, id)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func (c *commentRepository) EditReplyContent(ctx context.Context, id uint64, data CtrEditReply) error {
	query := `UPDATE tbl_replies SET content = $1 WHERE id = $2`
	_, err := c.db.ExecContext(ctx, query, data.Content, id)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func (c *commentRepository) EditReplyScore(ctx context.Context, id uint64, data CtrEditScore) error {
	query := `UPDATE tbl_replies SET score = $1 WHERE id = $2`
	_, err := c.db.ExecContext(ctx, query, data.Score, id)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func (c *commentRepository) EditCommentScore(ctx context.Context, id uint64, data CtrEditScore) error {
	query := `UPDATE tbl_comments SET score = $1 WHERE id = $2`
	_, err := c.db.ExecContext(ctx, query, data.Score, id)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}
