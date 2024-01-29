package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"main.go/api/models"
)

// CreateCategory godoc
// @Router /category [POST]
// @Summary Creates a new category
// @Description Creates a new category
// @Tags category
// @Accept json
// @Produce json
// @Param category body models.CreateCategory true "category"
// @Success      201  {object}  models.Response
// @Failure      400  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h Handler) CreateCategory(c *gin.Context) {
	category := models.CreateCategory{}

	if err := c.ShouldBindJSON(&category); err != nil {
		handleResponse(c, "error is while rading body from client", http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.storage.Category().Create(category)
	if err != nil {
		handleResponse(c, "error is while creating category", http.StatusInternalServerError, err.Error())
		return
	}

	createdCategory, err := h.storage.Category().GetByID(models.PrimaryKey{ID: id})
	if err != nil {
		handleResponse(c, "error is while getting by id", http.StatusInternalServerError, err.Error())
		return
	}

	handleResponse(c, "", http.StatusCreated, createdCategory)
}

// GetCategory godoc
// @Router  /category/{id} [GET]
// @Summary Gets category
// @Description get catgory by id
// @Tags category
// @Accept json
// @Produce json
// @Param id path string true "category"
// @Success      200  {object}  models.Response
// @Failure      400  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h Handler) GetCategory(c *gin.Context) {
	uid := c.Param("id")

	category, err := h.storage.Category().GetByID(models.PrimaryKey{ID: uid})
	if err != nil {
		handleResponse(c, "eror is while getting by id", http.StatusInternalServerError, err.Error())
		return
	}

	handleResponse(c, "", http.StatusOK, category)
}

// GetCategoryList godoc
// @Router       /categories [GET]
// @Summary      Get category list
// @Description  get category list
// @Tags         category
// @Accept       json
// @Produce      json
// @Param        page query string false "page"
// @Param 		 limit query string false "limit"
// @Param 		 search query string false "search"
// @Success      200  {object}  models.CategoryResponse
// @Failure      400  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h Handler) GetCategoryList(c *gin.Context) {
	var (
		page, limit int
		search      string
		err         error
	)

	pageStr := c.DefaultQuery("page", "1")
	page, err = strconv.Atoi(pageStr)
	if err != nil {
		handleResponse(c, "error is while converting page", http.StatusBadRequest, err.Error())
		return
	}

	limitStr := c.DefaultQuery("limit", "10")
	limit, err = strconv.Atoi(limitStr)
	if err != nil {
		handleResponse(c, "error is while converting limit", http.StatusBadRequest, err.Error())
		return
	}

	search = c.Query("search")

	categories, err := h.storage.Category().GetList(models.GetListRequest{
		Page:   page,
		Limit:  limit,
		Search: search,
	})

	if err != nil {
		handleResponse(c, "error is while get list", http.StatusInternalServerError, err.Error())
		return
	}

	handleResponse(c, "", http.StatusOK, categories)
}

// UpdateCetgory godoc
// @Router       /category/{id} [PUT]
// @Summary      Update category
// @Description  update category
// @Tags         category
// @Accept       json
// @Produce      json
// @Param 		 id path string true "category_id"
// @Param        user body models.UpdateCategory true "category"
// @Success      200  {object}  models.Category
// @Failure      400  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h Handler) UpdateCategory(c *gin.Context) {
	category := models.UpdateCategory{}
	uid := c.Param("id")

	if err := c.ShouldBindJSON(&category); err != nil {
		handleResponse(c, "error is while reading body", http.StatusBadRequest, err.Error())
		return
	}

	category.ID = uid

	id, err := h.storage.Category().Update(category)
	if err != nil {
		handleResponse(c, "error is while updating", http.StatusInternalServerError, err.Error())
		return
	}

	updatedCategory, err := h.storage.Category().GetByID(models.PrimaryKey{ID: id})
	if err != nil {
		handleResponse(c, "error is while getting by id", http.StatusInternalServerError, err.Error())
		return
	}

	handleResponse(c, "", http.StatusOK, updatedCategory)
}

// DeleteCategory godoc
// @Router       /category/{id} [DELETE]
// @Summary      Delete category
// @Description  delete category
// @Tags         category
// @Accept       json
// @Produce      json
// @Param 		 id path string true "category_id"
// @Param        category body models.UpdateCategory true "category"
// @Success      200  {object}  models.Response
// @Failure      400  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h Handler) DeleteCategory(c *gin.Context) {
	uid := c.Param("id")

	if err := h.storage.Category().Delete(models.PrimaryKey{ID: uid}); err != nil {
		handleResponse(c, "error is while delete", http.StatusInternalServerError, err.Error())
		return
	}

	handleResponse(c, "", http.StatusOK, "category deleted!")
}
