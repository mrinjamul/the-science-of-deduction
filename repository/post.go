package repository

import (
	"github.com/gin-gonic/gin"
	"github.com/mrinjamul/the-science-of-deduction/models"
	"gorm.io/gorm"
)

// CaseFiles is the repository for CaseFiles
type CaseFiles interface {
	// CreateCaseFile creates a new CaseFile
	CreateCaseFile(ctx *gin.Context, caseFile *models.CaseFiles) error
	// GetCaseFile returns a CaseFile
	GetCaseFileByID(ctx *gin.Context, id int) (models.CaseFiles, error)
	// GetCaseFileByTitle returns a CaseFile
	GetCaseFileByTitle(ctx *gin.Context, title string) (models.CaseFiles, error)
	// GetCaseFiles returns a list of CaseFiles
	GetCaseFiles(ctx *gin.Context) ([]models.CaseFiles, error)
	// GetArchivedCaseFiles returns a list of archived CaseFiles
	GetArchivedCaseFiles(ctx *gin.Context) ([]models.CaseFiles, error)
	// GetActiveCaseFiles returns a list of active CaseFiles
	GetActiveCaseFiles(ctx *gin.Context) ([]models.CaseFiles, error)
	// GetOnGoingCaseFiles returns a list of ongoing CaseFiles
	GetOnGoingCaseFiles(ctx *gin.Context) ([]models.CaseFiles, error)
	// GetRecentCaseFiles returns a list of recent CaseFiles
	GetRecentCaseFiles(ctx *gin.Context) ([]models.CaseFiles, error)
	// UpdateCaseFile updates a CaseFile
	UpdateCaseFile(ctx *gin.Context, caseFile *models.CaseFiles) error
	// DeleteCaseFile deletes a CaseFile
	DeleteCaseFile(ctx *gin.Context, id int) error
}

type post struct {
	db gorm.DB
}

// CreateCaseFile creates a new CaseFile
func (p *post) CreateCaseFile(ctx *gin.Context, caseFile *models.CaseFiles) error {
	return p.db.Create(caseFile).Error
}

// GetCaseFile returns a CaseFile
func (p *post) GetCaseFileByID(ctx *gin.Context, id int) (models.CaseFiles, error) {
	var caseFile models.CaseFiles
	err := p.db.First(&caseFile, uint(id)).Error
	return caseFile, err
}

// GetCaseFileByTitle returns a CaseFile
func (p *post) GetCaseFileByTitle(ctx *gin.Context, title string) (models.CaseFiles, error) {
	var caseFile models.CaseFiles
	err := p.db.Where("title = ?", title).First(&caseFile).Error
	return caseFile, err
}

// GetCaseFiles returns a list of CaseFiles
func (p *post) GetCaseFiles(ctx *gin.Context) ([]models.CaseFiles, error) {
	var caseFiles []models.CaseFiles
	err := p.db.Order("created_at desc").Find(&caseFiles).Error
	return caseFiles, err
}

// GetArchivedCaseFiles returns a list of archived CaseFiles
func (p *post) GetArchivedCaseFiles(ctx *gin.Context) ([]models.CaseFiles, error) {
	var caseFiles []models.CaseFiles
	err := p.db.Order("created_at desc").Where("is_archived = ?", 1).Find(&caseFiles).Error
	return caseFiles, err
}

// GetActiveCaseFiles returns a list of active CaseFiles
func (p *post) GetActiveCaseFiles(ctx *gin.Context) ([]models.CaseFiles, error) {
	var caseFiles []models.CaseFiles
	// case files which are not archived and closed
	err := p.db.Order("created_at desc").Where("is_archived = ? AND is_closed = ?", 0, 1).Find(&caseFiles).Error
	return caseFiles, err
}

// GetOnGoingCaseFiles returns a list of ongoing CaseFiles
func (p *post) GetOnGoingCaseFiles(ctx *gin.Context) ([]models.CaseFiles, error) {
	var caseFiles []models.CaseFiles
	// order by create time and not closed and archived
	err := p.db.Order("created_at desc").Where("is_closed = ? AND is_archived = ?", 0, 0).Find(&caseFiles).Error
	return caseFiles, err
}

// GetRecentCaseFiles returns a list of recent CaseFiles
func (p *post) GetRecentCaseFiles(ctx *gin.Context) ([]models.CaseFiles, error) {
	var caseFiles []models.CaseFiles
	err := p.db.Order("created_at desc").Limit(5).Find(&caseFiles).Error
	return caseFiles, err
}

// UpdateCaseFile updates a CaseFile
func (p *post) UpdateCaseFile(ctx *gin.Context, caseFile *models.CaseFiles) error {
	// update the caseFile
	return p.db.Save(caseFile).Error
}

// DeleteCaseFile deletes a CaseFile
func (p *post) DeleteCaseFile(ctx *gin.Context, id int) error {
	// fetch the caseFile
	var caseFile models.CaseFiles
	err := p.db.First(&caseFile, id).Error
	if err != nil {
		return err
	}
	// delete the caseFile
	return p.db.Delete(&caseFile).Error
}

// NewCaseFiles returns a new CaseFiles repository
func NewCaseFiles(db *gorm.DB) CaseFiles {
	return &post{
		db: *db,
	}
}
