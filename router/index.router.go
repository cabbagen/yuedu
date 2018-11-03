package router

import (
  "yuedu/controller"
  "github.com/gin-gonic/gin"
)

var indexController controller.IndexController;

var indexRouter []descriptor = []descriptor {
  descriptor {
    path: "/index",
    method: "GET",
    handlers:  []gin.HandlerFunc{ indexController.HandleIndex },
  },
}
