package controller

type User struct {
	// Add Fields
}

// ref: https://github.com/swaggo/swag/blob/master/example/celler/controller/accounts.go
type UserCtrl struct{}

// @Summary Get all items
// @Description Get all items
// @Accept  json
// @Produce  json
// @Success 200 {object} []model.User
// @Router /User [get]
func (ctrl *UserCtrl) Get(c *gin.Context) {

	// c.JSON(http.StatusOK, gin.H{})
	util.Json(http.StatusOK, "success", {"name": "test"})
}

// @Summary Get one item
// @Description Get one item
// @Accept  json
// @Produce  json
// @Param   id     path    int     true        "ID"
// @Success 200 {object} model.User
// @Failure 400 {string} string "400 StatusBadRequest"
// @Failure 404 {string} string "404 not found"
// @Router /User/{id} [get]
func (ctrl *UserCtrl) GetOne(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}

// @Summary Add item
// @Description Add item
// @Accept  json
// @Produce  json
// @Param role body model.Usertrue "data"
// @Success 200 {object} model.User
// @Router /User [post]
func (ctrl *UserCtrl) Post(c *gin.Context) {
	var t User
	if err := c.ShouldBindJSON(&t); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}

// @Summary Update item
// @Description Update item
// @Accept  json
// @Produce  json
// @Param role body model.User true "data"
// @Param   id     path    int     true        "ID"
// @Success 200 {object} model.User
// @Failure 400 {string} string "400 StatusBadRequest"
// @Failure 404 {string} string "404 not found"
// @Router /User/{id} [put]
func (ctrl *UserCtrl) Put(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	var t User
	if err := c.ShouldBindJSON(&t); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}

// @Summary Delete item
// @Description Delete item
// @Accept  json
// @Produce  json
// @Param   id     path    int     true        "ID"
// @Success 204
// @Failure 400 {string} string "400 StatusBadRequest"
// @Failure 404 {string} string "404 not found"
// @Router /User/{id} [delete]
func (ctrl *UserCtrl) Delete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusNoContent, gin.H{})
}
