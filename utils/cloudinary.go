package utils

import (
	"context"
	"fmt"
	"mime/multipart"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
)

// CloudinaryService handles Cloudinary operations
type CloudinaryService struct {
	cld *cloudinary.Cloudinary
}

// NewCloudinaryService creates a new Cloudinary service instance
func NewCloudinaryService() (*CloudinaryService, error) {
	cloudName := os.Getenv("CLOUDINARY_CLOUD_NAME")
	apiKey := os.Getenv("CLOUDINARY_API_KEY")
	apiSecret := os.Getenv("CLOUDINARY_API_SECRET")

	if cloudName == "" || apiKey == "" || apiSecret == "" {
		return nil, fmt.Errorf("cloudinary credentials not properly configured")
	}

	cld, err := cloudinary.NewFromParams(cloudName, apiKey, apiSecret)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize cloudinary: %v", err)
	}

	return &CloudinaryService{cld: cld}, nil
}

// UploadProfileImage uploads a profile image to Cloudinary
func (cs *CloudinaryService) UploadProfileImage(file *multipart.FileHeader, userID string) (string, error) {
	// Open the file
	src, err := file.Open()
	if err != nil {
		return "", fmt.Errorf("failed to open file: %v", err)
	}
	defer src.Close()

	// Validate file type
	if !isValidImageType(file.Filename) {
		return "", fmt.Errorf("invalid file type. Only JPG, JPEG, PNG, and GIF are allowed")
	}

	// Validate file size (max 5MB)
	if file.Size > 5*1024*1024 {
		return "", fmt.Errorf("file size too large. Maximum size is 5MB")
	}

	// Create a unique public ID
	timestamp := time.Now().UnixNano()
	publicID := fmt.Sprintf("profile_images/%s_%d", userID, timestamp)

	// Upload parameters
	uploadParams := uploader.UploadParams{
		PublicID:       publicID,
		Folder:         "meetup_profiles",
		ResourceType:   "image",
		Format:         "jpg", // Convert all images to JPG for consistency
		Transformation: "w_400,h_400,c_fill,g_face,q_auto", // Resize to 400x400, focus on face, auto quality
		Tags:           []string{"profile", "user_" + userID},
	}

	// Upload to Cloudinary
	ctx := context.Background()
	result, err := cs.cld.Upload.Upload(ctx, src, uploadParams)
	if err != nil {
		return "", fmt.Errorf("failed to upload to cloudinary: %v", err)
	}

	return result.SecureURL, nil
}

// DeleteProfileImage deletes a profile image from Cloudinary
func (cs *CloudinaryService) DeleteProfileImage(imageURL string) error {
	// Extract public ID from URL
	publicID := extractPublicIDFromURL(imageURL)
	if publicID == "" {
		return fmt.Errorf("invalid cloudinary URL")
	}

	ctx := context.Background()
	_, err := cs.cld.Upload.Destroy(ctx, uploader.DestroyParams{
		PublicID: publicID,
	})

	return err
}

// isValidImageType checks if the file has a valid image extension
func isValidImageType(filename string) bool {
	ext := strings.ToLower(filepath.Ext(filename))
	validTypes := []string{".jpg", ".jpeg", ".png", ".gif"}
	
	for _, validType := range validTypes {
		if ext == validType {
			return true
		}
	}
	return false
}

// extractPublicIDFromURL extracts the public ID from a Cloudinary URL
func extractPublicIDFromURL(url string) string {
	// Example URL: https://res.cloudinary.com/demo/image/upload/v1234567890/meetup_profiles/profile_images/user123_1234567890.jpg
	parts := strings.Split(url, "/")
	if len(parts) < 7 {
		return ""
	}

	// Find the upload part
	uploadIndex := -1
	for i, part := range parts {
		if part == "upload" {
			uploadIndex = i
			break
		}
	}

	if uploadIndex == -1 || uploadIndex+2 >= len(parts) {
		return ""
	}

	// Get everything after version (v1234567890)
	pathParts := parts[uploadIndex+2:]
	publicID := strings.Join(pathParts, "/")
	
	// Remove file extension
	if dotIndex := strings.LastIndex(publicID, "."); dotIndex != -1 {
		publicID = publicID[:dotIndex]
	}

	return publicID
}
