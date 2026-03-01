package dtos

import "time"

type PostResponseDTO struct {
	ID        uint      `json:"id"`
	Title     string    `json:"title"`
	Slug      string    `json:"slug"`
	Content   string    `json:"content"`
	Published bool      `json:"published"`
	CreatedAt time.Time `json:"created_at"`
	// Nested User menggunakan pointer agar bisa nil jika tidak di-Preload
	Author *UserResponseDTO `json:"author,omitempty"`
}
