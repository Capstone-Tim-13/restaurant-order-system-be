package repository

import (
	"capstone/features/menu"
	"context"
	"mime/multipart"
	"os"

	"github.com/cloudinary/cloudinary-go"
	"github.com/cloudinary/cloudinary-go/api/uploader"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type MenuRepositoryImpl struct {
	db  *gorm.DB
	cdn *cloudinary.Cloudinary
}

func NewMenuRepository(db *gorm.DB, cdn *cloudinary.Cloudinary) menu.Repository {
	return &MenuRepositoryImpl{db: db, cdn: cdn}
}

func (r *MenuRepositoryImpl) Save(Newmenu *menu.Menu) (*menu.Menu, error) {
	result := r.db.Create(&Newmenu)
	if result.Error != nil {
		return nil, result.Error
	}
	return Newmenu, nil
}

func (r *MenuRepositoryImpl) UploadImage(ctx context.Context, file multipart.File, name string) (string, error) {
	response, err := r.cdn.Upload.Upload(ctx, file, uploader.UploadParams{
		Folder:   os.Getenv("CDN_FOLDER_NAME"),
		PublicID: name,
	})
	if err != nil {
		logrus.Error("Repository: Upload image error,", err)
		return "", err
	}

	return response.SecureURL, nil

	// var file = helpers.OpenFileHeader(fileHeader)

	// ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	// defer cancel()

	// response, err := r.cdn.Upload.Upload(ctx, file, uploader.UploadParams{
	// 	Folder:   r.config.CDN_FOLDER_NAME,
	// 	PublicID: name,
	// })
	// if err != nil {
	// 	logrus.Error("Repository: Upload image error,", err)
	// 	return "", err
	// }

	// return response.SecureURL, nil
}

func (r *MenuRepositoryImpl) Update(Newmenu *menu.Menu) (*menu.Menu, error) {
	result := r.db.Table("menus").Where("id = ?", Newmenu.ID).Updates(menu.Menu{Name: Newmenu.Name, CategoryID: Newmenu.CategoryID, Description: Newmenu.Description, Price: Newmenu.Price})
	if result.Error != nil {
		return nil, result.Error
	}
	return Newmenu, nil
}

func (r *MenuRepositoryImpl) FindAll() ([]menu.Menu, error) {
	menu := []menu.Menu{}

	result := r.db.Where("delete_at IS NULL").Find(&menu)
	if result.Error != nil {
		return nil, result.Error
	}

	return menu, nil
}

func (r *MenuRepositoryImpl) FindById(id int) (*menu.Menu, error) {
	menu := menu.Menu{}

	result := r.db.Where("delete_at IS NULL").First(&menu, id)
	if result.Error != nil {
		return nil, result.Error
	}

	return &menu, nil
}

func (r *MenuRepositoryImpl) FindByName(name string) (*menu.Menu, error) {
	menu := menu.Menu{}

	result := r.db.Where("name = ?", name).First(&menu)
	if result.Error != nil {
		return nil, result.Error
	}

	return &menu, nil
}

func (r *MenuRepositoryImpl) FindByCategoryId(categoryId int) ([]menu.Menu, error) {
	menu := []menu.Menu{}

	result := r.db.Where("category_id = ?", categoryId).Find(&menu)
	if result.Error != nil {
		return nil, result.Error
	}

	return menu, nil
}

func (r *MenuRepositoryImpl) Delete(id int) error {
	result := r.db.Table("menus").Where("id = ?", id).Unscoped().Delete(id)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
