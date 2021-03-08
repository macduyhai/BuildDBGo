package controllers

import (
	"BuildDBGo/dtos"
	"BuildDBGo/services"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"strings"
	"time"

	"BuildDBGo/utils"

	"github.com/gin-gonic/gin"
)

const TimeFormatFull = "2006-01-02 15:04:05"

type Controller struct {
	OrderService       services.OrderService
	BranchSellService  services.BranchSellService
	TypeProductService services.TypeProductService
	SellerService      services.SellerService
	AuthenService      services.AuthenService
}

func (c Controller) HealthCheck(contex *gin.Context) {
	contex.JSON(200, gin.H{
		"status": "running",
	})
}

// ----- Delete method DeleteBranchSell

func (c Controller) DeleteSeller(ctx *gin.Context) {
	var request dtos.DeleteSellerRequest
	bytes, err := ioutil.ReadAll(ctx.Request.Body)
	if err != nil {
		fmt.Println("get raw body error", err)
		utils.ResponseErrorGin(ctx, "get raw body error")
		return
	}

	err = json.Unmarshal(bytes, &request)
	if err != nil {
		fmt.Println("bind json error", err, "raw_body", string(bytes))
		utils.ResponseErrorGin(ctx, "bind json error")
		return
	}
	log.Println(request.Name)
	c.SellerService.Detete(request.Name)
	// err := c.BranchSellService.Detete(request.Name)
	// if err != nil {
	// 	fmt.Println("delete branch error", err)
	// 	utils.ResponseErrorGin(ctx, "delete branch error")
	// 	return
	// }

	fmt.Println("Delete type product success")
	utils.ResponseSuccess(ctx, nil)
}

func (c Controller) DeleteTypeProduct(ctx *gin.Context) {
	var request dtos.DeleteTypeRequest
	bytes, err := ioutil.ReadAll(ctx.Request.Body)
	if err != nil {
		fmt.Println("get raw body error", err)
		utils.ResponseErrorGin(ctx, "get raw body error")
		return
	}

	err = json.Unmarshal(bytes, &request)
	if err != nil {
		fmt.Println("bind json error", err, "raw_body", string(bytes))
		utils.ResponseErrorGin(ctx, "bind json error")
		return
	}
	log.Println(request.Name)
	c.TypeProductService.Detete(request.Name)
	// err := c.BranchSellService.Detete(request.Name)
	// if err != nil {
	// 	fmt.Println("delete branch error", err)
	// 	utils.ResponseErrorGin(ctx, "delete branch error")
	// 	return
	// }

	fmt.Println("Delete type product success")
	utils.ResponseSuccess(ctx, nil)
}

func (c Controller) DeleteBranchSell(ctx *gin.Context) {
	var request dtos.DeleteBranchRequest
	bytes, err := ioutil.ReadAll(ctx.Request.Body)
	if err != nil {
		fmt.Println("get raw body error", err)
		utils.ResponseErrorGin(ctx, "get raw body error")
		return
	}

	err = json.Unmarshal(bytes, &request)
	if err != nil {
		fmt.Println("bind json error", err, "raw_body", string(bytes))
		utils.ResponseErrorGin(ctx, "bind json error")
		return
	}
	log.Println(request.Name)
	c.BranchSellService.Detete(request.Name)
	// err := c.BranchSellService.Detete(request.Name)
	// if err != nil {
	// 	fmt.Println("delete branch error", err)
	// 	utils.ResponseErrorGin(ctx, "delete branch error")
	// 	return
	// }

	fmt.Println("Delete branch success")
	utils.ResponseSuccess(ctx, nil)
}

// ----- ADD Authen
func (c Controller) AddAuthen(ctx *gin.Context) {

	// log.Println(ctx.Query("Key"))
	// log.Println(ctx.Query("CreatedAt"))

	var request dtos.AuthenKey
	// log.Println(ctx.Request.Body)
	// bytes, err := ioutil.ReadAll(ctx.Request.Body)
	// if err != nil {
	// 	fmt.Println("get raw body error", err)
	// 	utils.ResponseErrorGin(ctx, "get raw body error")
	// 	return
	// }
	// log.Println(bytes)

	// err = json.Unmarshal(bytes, &request)
	// if err != nil {
	// 	fmt.Println("bind json error", err, "raw_body", string(bytes))
	// 	utils.ResponseErrorGin(ctx, "bind json error")
	// 	return
	// }
	log.Println(request)
	request.Key = ctx.Query("Key")
	// request.CreatedAt = ctx.Query("CreatedAt")

	t := time.Now()
	t = t.Add(time.Hour * 7)
	request.CreatedAt = &t

	log.Println(request)
	resp, err := c.AuthenService.AddAuthen(request)
	if err != nil {
		fmt.Println("add order error", err)
		utils.ResponseErrorGin(ctx, "add order error")
		return
	}

	fmt.Println("add success")
	utils.ResponseSuccess(ctx, resp)
}

// ----- ADD METHOD
func (c Controller) AddSeller(ctx *gin.Context) {
	var request dtos.AddsellerRequest
	bytes, err := ioutil.ReadAll(ctx.Request.Body)
	if err != nil {
		fmt.Println("get raw body error", err)
		utils.ResponseErrorGin(ctx, "get raw body error")
		return
	}

	err = json.Unmarshal(bytes, &request)
	if err != nil {
		fmt.Println("bind json error", err, "raw_body", string(bytes))
		utils.ResponseErrorGin(ctx, "bind json error")
		return
	}
	// request.Sellers[0].Name = strings.ToLower(request.Sellers.Name)
	resp, err := c.SellerService.AddSeller(request)
	if err != nil {
		fmt.Println("add order error", err)
		utils.ResponseErrorGin(ctx, "add order error")
		return
	}

	fmt.Println("add success")
	utils.ResponseSuccess(ctx, resp)
}

func (c Controller) AddTypeProduct(ctx *gin.Context) {
	var request dtos.AddtypeRequest
	bytes, err := ioutil.ReadAll(ctx.Request.Body)
	if err != nil {
		fmt.Println("get raw body error", err)
		utils.ResponseErrorGin(ctx, "get raw body error")
		return
	}

	err = json.Unmarshal(bytes, &request)
	if err != nil {
		fmt.Println("bind json error", err, "raw_body", string(bytes))
		utils.ResponseErrorGin(ctx, "bind json error")
		return
	}
	// request.TypeProducts[0].Name = strings.ToLower(request.TypeProducts.Name)
	resp, err := c.TypeProductService.AddTypeProduct(request)
	if err != nil {
		fmt.Println("add order error", err)
		utils.ResponseErrorGin(ctx, "add order error")
		return
	}

	fmt.Println("add success")
	utils.ResponseSuccess(ctx, resp)
}

func (c Controller) AddBranchSell(ctx *gin.Context) {
	var request dtos.AddbranchRequest
	bytes, err := ioutil.ReadAll(ctx.Request.Body)
	if err != nil {
		fmt.Println("get raw body error", err)
		utils.ResponseErrorGin(ctx, "get raw body error")
		return
	}

	err = json.Unmarshal(bytes, &request)
	if err != nil {
		fmt.Println("bind json error", err, "raw_body", string(bytes))
		utils.ResponseErrorGin(ctx, "bind json error")
		return
	}
	// request.BranchSells[0].Name = strings.ToLower(request.BranchSells.Name)
	resp, err := c.BranchSellService.AddBranchSell(request)
	if err != nil {
		fmt.Println("add order error", err)
		utils.ResponseErrorGin(ctx, "add order error")
		return
	}

	fmt.Println("add success")
	utils.ResponseSuccess(ctx, resp)
}

func (c Controller) AddOrder(ctx *gin.Context) {
	var request dtos.AddOrderRequest

	//  var request dtos.AddOrderNewRequest
	bytes, err := ioutil.ReadAll(ctx.Request.Body)
	if err != nil {
		fmt.Println("get raw body ADD ORDER error", err)
		utils.ResponseErrorGin(ctx, "get raw body error")
		return
	}

	err = json.Unmarshal(bytes, &request)
	if err != nil {
		fmt.Println("bind json ADD ORDER error", err, "raw_body", string(bytes))
		utils.ResponseErrorGin(ctx, "bind json error")
		return
	}
	resp, err := c.OrderService.AddOrder(request)
	if err != nil {
		fmt.Println("add order error", err)
		utils.ResponseErrorGin(ctx, "add order error")
		return
	}

	fmt.Println("add success")
	utils.ResponseSuccess(ctx, resp)
}

// AddFullOrder
func (c Controller) AddFullOrder(ctx *gin.Context) {
	var request dtos.AddfullOrderRequest

	//  var request dtos.AddOrderNewRequest
	bytes, err := ioutil.ReadAll(ctx.Request.Body)
	if err != nil {
		fmt.Println("get raw body ADD FULL ORDER From tools error", err)
		utils.ResponseErrorGin(ctx, "get raw body error")
		return
	}

	err = json.Unmarshal(bytes, &request)
	if err != nil {
		log.Println("bind json ADD FULL ORDER error", err, "raw_body", string(bytes))
		utils.ResponseErrorGin(ctx, "bind json error")
		return
	}
	resp, err := c.OrderService.AddFullOrder(request)
	if err != nil {
		log.Println("add order error", err)
		utils.ResponseErrorGin(ctx, err.Error())
		return
	}

	fmt.Println("add success")
	utils.ResponseSuccess(ctx, resp)
}

// UpdateSeller

func (c Controller) UpdateSeller(ctx *gin.Context) {
	var request dtos.Seller
	bytes, err := ioutil.ReadAll(ctx.Request.Body)
	if err != nil {
		fmt.Println("get raw body error", err)
		utils.ResponseErrorGin(ctx, "get raw body error")
		return
	}

	err = json.Unmarshal(bytes, &request)
	if err != nil {
		fmt.Println("bind json error", err, "raw_body", string(bytes))
		utils.ResponseErrorGin(ctx, "bind json error")
		return
	}
	request.Name = strings.ToLower(request.Name)
	if request.Name == "" {
		fmt.Println("required Name seller", "raw_body", string(bytes))
		utils.ResponseErrorGin(ctx, "Name Seller")
		return
	}

	resp, err := c.SellerService.Updates(request)
	if err != nil {
		fmt.Println("updates seller error", err)
		utils.ResponseErrorGin(ctx, "update seller error")
		return
	}

	fmt.Println("updates seller success", "ID", resp.ID)
	utils.ResponseSuccess(ctx, resp)
}

// UpdateTypeProduct
func (c Controller) UpdateTypeProduct(ctx *gin.Context) {
	var request dtos.TypeProduct
	bytes, err := ioutil.ReadAll(ctx.Request.Body)
	if err != nil {
		fmt.Println("get raw body error", err)
		utils.ResponseErrorGin(ctx, "get raw body error")
		return
	}

	err = json.Unmarshal(bytes, &request)
	if err != nil {
		fmt.Println("bind json error", err, "raw_body", string(bytes))
		utils.ResponseErrorGin(ctx, "bind json error")
		return
	}
	request.Name = strings.ToLower(request.Name)
	if request.Name == "" {
		fmt.Println("required Name Type Product", "raw_body", string(bytes))
		utils.ResponseErrorGin(ctx, "Name Type Product")
		return
	}

	resp, err := c.TypeProductService.Updates(request)
	if err != nil {
		fmt.Println("updates Type Product error", err)
		utils.ResponseErrorGin(ctx, "update branch error")
		return
	}

	fmt.Println("updates Type Product success", "ID", resp.ID)
	utils.ResponseSuccess(ctx, resp)
}

func (c Controller) UpdateBranch(ctx *gin.Context) {
	var request dtos.BranchSell
	bytes, err := ioutil.ReadAll(ctx.Request.Body)
	if err != nil {
		fmt.Println("get raw body error", err)
		utils.ResponseErrorGin(ctx, "get raw body error")
		return
	}

	err = json.Unmarshal(bytes, &request)
	if err != nil {
		fmt.Println("bind json error", err, "raw_body", string(bytes))
		utils.ResponseErrorGin(ctx, "bind json error")
		return
	}
	request.Name = strings.ToLower(request.Name)
	if request.Name == "" {
		fmt.Println("required Name branch Sell", "raw_body", string(bytes))
		utils.ResponseErrorGin(ctx, "Name branch Sell")
		return
	}

	resp, err := c.BranchSellService.Updates(request)
	if err != nil {
		fmt.Println("updates branch error", err)
		utils.ResponseErrorGin(ctx, "update branch error")
		return
	}

	fmt.Println("updates branch success", "ID", resp.ID)
	utils.ResponseSuccess(ctx, resp)
}
func (c Controller) Printers(ctx *gin.Context) {
	var request dtos.PrintersRequest
	bytes, err := ioutil.ReadAll(ctx.Request.Body)
	if err != nil {
		fmt.Println("get raw body error", err)
		utils.ResponseErrorGin(ctx, "get raw body error")
		return
	}

	err = json.Unmarshal(bytes, &request)
	if err != nil {
		fmt.Println("bind json error", err, "raw_body", string(bytes))
		utils.ResponseErrorGin(ctx, "bind json error")
		return
	}

	// if request.OrderNumber == "" {
	// 	fmt.Println("required OrderNumber", "raw_body", string(bytes))
	// 	utils.ResponseErrorGin(ctx, "required OrderNumber")
	// 	return
	// }

	resp, err := c.OrderService.Printers(request)
	if err != nil {
		fmt.Println("updates order error", err)
		utils.ResponseErrorGin(ctx, "update order error")
		return
	}

	fmt.Println("updates order success", "ID", resp.ID)
	utils.ResponseSuccess(ctx, resp)
}
func (c Controller) UpdateOrders(ctx *gin.Context) {
	var request dtos.Order
	bytes, err := ioutil.ReadAll(ctx.Request.Body)
	if err != nil {
		fmt.Println("get raw body error", err)
		utils.ResponseErrorGin(ctx, "get raw body error")
		return
	}

	err = json.Unmarshal(bytes, &request)
	if err != nil {
		fmt.Println("bind json error", err, "raw_body", string(bytes))
		utils.ResponseErrorGin(ctx, "bind json error")
		return
	}

	if request.OrderNumber == "" {
		fmt.Println("required OrderNumber", "raw_body", string(bytes))
		utils.ResponseErrorGin(ctx, "required OrderNumber")
		return
	}

	resp, err := c.OrderService.Updates(request)
	if err != nil {
		fmt.Println("updates order error", err)
		utils.ResponseErrorGin(ctx, "update order error")
		return
	}

	fmt.Println("updates order success", "ID", resp.ID)
	utils.ResponseSuccess(ctx, resp)
}

func (c Controller) AddLabelToOrder(ctx *gin.Context) {
	var request dtos.AddLabelRequest
	bytes, err := ioutil.ReadAll(ctx.Request.Body)
	if err != nil {
		log.Println("get raw body AddLabelToOrder error", err)
		utils.ResponseErrorGin(ctx, "get raw body error")
		return
	}
	err = json.Unmarshal([]byte(bytes), &request)
	if err != nil {
		log.Println("bind json AddLabelToOrder error", err, "raw_body", string(bytes))
		utils.ResponseErrorGin(ctx, "bind json error")
		return
	}
	// log.Println(request)
	// if request.Items == nil {
	if request.LableDetails.PartnerTrackingNumber == "" ||
		request.LableDetails.TrackingNumber == "" ||
		request.LableDetails.URL == "" {
		if request.Items == nil {
			log.Println("require field in label details is missing", request)
			utils.ResponseErrorGin(ctx, "require field is missing")
			return
		} else {
			res, err := c.OrderService.AddLabelsToItems(request)
			if err != nil {
				log.Println("add labels to order error", err)
				fmt.Println(res)
				utils.ResponseErrorGin(ctx, "add Item to items error")
				return
			} else {
			}

			fmt.Println("add labels to order done")
			utils.ResponseSuccess(ctx, res)
		}
	} else {
		res, err := c.OrderService.AddLabelsToOrder(request)
		if err != nil {
			log.Println("add labels to order error", err)
			utils.ResponseErrorGin(ctx, "add labels to order error")
			return
		} else {
			// fmt.Println()
		}

		utils.ResponseSuccess(ctx, res)
	}

	// }

	// // AddLabelsToItems

	// res, err := c.OrderService.AddLabelsToItems(request)
	// if err != nil {
	// 	// fmt.Println("add labels to order error", err)
	// 	fmt.Println(res)
	// 	utils.ResponseErrorGin(ctx, "add Item to items error")
	// 	return
	// } else {
	// }

	// fmt.Println("add labels to order done")
	// utils.ResponseSuccess(ctx, res)
}

// Search Seller
func (c Controller) getSearchQuerySeller(ctx *gin.Context) ([]dtos.SearchSellerQuery, error) {
	result := make([]dtos.SearchSellerQuery, 0)
	result = append(result, dtos.SearchSellerQuery{
		Key:   "deleted_at IS NULL",
		Value: nil,
	})

	begin := ctx.Query("begin_time")
	if begin != "" {
		item := dtos.SearchSellerQuery{
			Key:   "created_at > ?",
			Value: begin,
		}
		result = append(result, item)
	}

	end := ctx.Query("end_time")
	if end != "" {
		item := dtos.SearchSellerQuery{
			Key:   "created_at < ?",
			Value: end,
		}
		result = append(result, item)
	}

	name := ctx.Query("name")
	if name != "" {
		item := dtos.SearchSellerQuery{
			Key:   "name = ?",
			Value: name,
		}
		result = append(result, item)
	}

	return result, nil
}

//Search Authen Key
func (c Controller) SearchAuthen(ctx *gin.Context) {

	resp, err := c.AuthenService.SearchAuthen()
	if err != nil {
		fmt.Println("search orders error", err)
		utils.ResponseErrorGin(ctx, "search order error")
		return
	}

	fmt.Println("search success")
	utils.ResponseSuccess(ctx, resp)
}

func (c Controller) SearchSeller(ctx *gin.Context) {
	queries, err := c.getSearchQuerySeller(ctx)
	if err != nil {
		fmt.Println("bind json error", err)
		utils.ResponseErrorGin(ctx, "bind json error")
		return
	}

	resp, err := c.SellerService.SearchSeller(queries)
	if err != nil {
		fmt.Println("search orders error", err)
		utils.ResponseErrorGin(ctx, "search order error")
		return
	}

	fmt.Println("search success")
	utils.ResponseSuccess(ctx, resp)
}

// Search Type Product
func (c Controller) getSearchQueryType(ctx *gin.Context) ([]dtos.SearchTypeProductQuery, error) {
	result := make([]dtos.SearchTypeProductQuery, 0)
	result = append(result, dtos.SearchTypeProductQuery{
		Key:   "deleted_at IS NULL",
		Value: nil,
	})

	begin := ctx.Query("begin_time")
	if begin != "" {
		item := dtos.SearchTypeProductQuery{
			Key:   "created_at > ?",
			Value: begin,
		}
		result = append(result, item)
	}

	end := ctx.Query("end_time")
	if end != "" {
		item := dtos.SearchTypeProductQuery{
			Key:   "created_at < ?",
			Value: end,
		}
		result = append(result, item)
	}

	name := ctx.Query("name")
	if name != "" {
		item := dtos.SearchTypeProductQuery{
			Key:   "name = ?",
			Value: name,
		}
		result = append(result, item)
	}

	return result, nil
}

// Search Branch Sell
func (c Controller) getSearchQueryBranch(ctx *gin.Context) ([]dtos.SearchBranchSellQuery, error) {
	result := make([]dtos.SearchBranchSellQuery, 0)
	result = append(result, dtos.SearchBranchSellQuery{
		Key:   "deleted_at IS NULL",
		Value: nil,
	})

	begin := ctx.Query("begin_time")
	if begin != "" {
		item := dtos.SearchBranchSellQuery{
			Key:   "created_at > ?",
			Value: begin,
		}
		result = append(result, item)
	}

	end := ctx.Query("end_time")
	if end != "" {
		item := dtos.SearchBranchSellQuery{
			Key:   "created_at < ?",
			Value: end,
		}
		result = append(result, item)
	}

	name := ctx.Query("name")
	if name != "" {
		item := dtos.SearchBranchSellQuery{
			Key:   "name = ?",
			Value: name,
		}
		result = append(result, item)
	}

	return result, nil
}

func (c Controller) SearchType(ctx *gin.Context) {
	queries, err := c.getSearchQueryType(ctx)
	if err != nil {
		fmt.Println("bind json error", err)
		utils.ResponseErrorGin(ctx, "bind json error")
		return
	}

	resp, err := c.TypeProductService.SearchType(queries)
	if err != nil {
		fmt.Println("search orders error", err)
		utils.ResponseErrorGin(ctx, "search order error")
		return
	}

	fmt.Println("search success")
	utils.ResponseSuccess(ctx, resp)
}

func (c Controller) SearchBranch(ctx *gin.Context) {
	queries, err := c.getSearchQueryBranch(ctx)
	if err != nil {
		fmt.Println("bind json error", err)
		utils.ResponseErrorGin(ctx, "bind json error")
		return
	}

	resp, err := c.BranchSellService.SearchBranch(queries)
	if err != nil {
		fmt.Println("search orders error", err)
		utils.ResponseErrorGin(ctx, "search order error")
		return
	}

	fmt.Println("search success")
	utils.ResponseSuccess(ctx, resp)
}

// Search Order
func (c Controller) getSearchQuery(ctx *gin.Context) ([]dtos.SearchQuery, error) {
	result := make([]dtos.SearchQuery, 0)
	result = append(result, dtos.SearchQuery{
		Key:   "deleted_at IS NULL",
		Value: nil,
	})

	begin := ctx.Query("begin_time")
	if begin != "" {
		item := dtos.SearchQuery{
			Key:   "created_at > ?",
			Value: begin,
		}
		result = append(result, item)
	}

	end := ctx.Query("end_time")
	if end != "" {
		item := dtos.SearchQuery{
			Key:   "created_at < ?",
			Value: end,
		}
		result = append(result, item)
	}

	// orderNumber := ctx.Query("order_number")
	// if orderNumber != "" {
	// 	item := dtos.SearchQuery{
	// 		Key:   "order_number = ?",
	// 		Value: orderNumber,
	// 	}
	// 	result = append(result, item)
	// }
	orderNumber := ctx.Query("order_number")
	if orderNumber != "" {
		item := dtos.SearchQuery{
			Key:   "order_number LIKE ?",
			Value: "%" + orderNumber + "%",
		}
		result = append(result, item)
	}

	status := ctx.Query("status")
	if status != "" {
		item := dtos.SearchQuery{
			Key:   "status=?",
			Value: status,
		}
		result = append(result, item)
	}

	checkid := ctx.Query("partner_tracking_number")
	if checkid != "" {
		item := dtos.SearchQuery{
			Key:   "partner_tracking_number=?",
			Value: checkid,
		}
		result = append(result, item)
	}

	return result, nil
}

const CommonTimeFormat = "2006-01-02"

// Check number order for week, month, year
func (c Controller) getOrderComplatedQuery(ctx *gin.Context, time_s time.Time, typ string, bra string, sel string, steptime string) ([]dtos.SearchQuery, error) {
	result := make([]dtos.SearchQuery, 0)
	var y int
	var m time.Month

	result = append(result, dtos.SearchQuery{
		Key:   "deleted_at IS NULL",
		Value: nil,
	})
	if steptime == "week" || steptime == "month" {
		// time_e := time_s.AddDate(0, 0, -1)
		log.Println(time_s.Format(CommonTimeFormat))
		// log.Println(time_s.Format(CommonTimeFormat) + " 23:59:59")
		item_start := dtos.SearchQuery{
			Key:   "time_completed > ?",
			Value: time_s.Format(CommonTimeFormat) + " 00:00:00",
		}
		result = append(result, item_start)

		item_end := dtos.SearchQuery{
			Key:   "time_completed < ?",
			Value: time_s.Format(CommonTimeFormat) + " 23:59:59",
		}
		result = append(result, item_end)
	} else if steptime == "year" {
		y, m, _ = time_s.Date()
		firstDay := time.Date(y, m, 1, 0, 0, 0, 0, time.UTC)
		lastDay := time.Date(y, m+1, 1, 0, 0, 0, -1, time.UTC)
		// time_e := time_s.AddDate(-1, 0, 0)
		// log.Println(firstDay.Format(CommonTimeFormat) + " 00:00:00")
		// log.Println(lastDay.Format(CommonTimeFormat) + " 23:59:59")
		item_start := dtos.SearchQuery{
			Key:   "time_completed > ?",
			Value: firstDay.Format(CommonTimeFormat) + " 00:00:00",
		}
		result = append(result, item_start)

		item_end := dtos.SearchQuery{
			Key:   "time_completed < ?",
			Value: lastDay.Format(CommonTimeFormat) + " 23:59:59",
		}
		result = append(result, item_end)
	} else {

	}
	status := "3"
	if status != "" {
		item := dtos.SearchQuery{
			Key:   "status=?",
			Value: status,
		}
		result = append(result, item)
	}
	if typ != "" {
		item := dtos.SearchQuery{
			Key:   "typeproduct=?",
			Value: typ,
		}
		result = append(result, item)
	}
	if bra != "" {
		item := dtos.SearchQuery{
			Key:   "branchsell=?",
			Value: bra,
		}
		result = append(result, item)
	}
	if sel != "" {
		item := dtos.SearchQuery{
			Key:   "seller=?",
			Value: sel,
		}
		result = append(result, item)
	}

	return result, nil
}
func (c Controller) getBranchnName(ctx *gin.Context) ([]dtos.BranchSell, error) {
	queries, err := c.getSearchQueryBranch(ctx)
	if err != nil {
		fmt.Println("bind json error", err)
		utils.ResponseErrorGin(ctx, "bind json error")
		return nil, err
	}
	resp, err := c.BranchSellService.SearchBranch(queries)
	return resp, nil
}
func (c Controller) getTypeName(ctx *gin.Context) ([]dtos.TypeProduct, error) {

	queries, err := c.getSearchQueryType(ctx)
	if err != nil {
		fmt.Println("bind json error", err)
		utils.ResponseErrorGin(ctx, "bind json error")
		return nil, err
	}

	resp, err := c.TypeProductService.SearchType(queries)
	return resp, nil
}
func (c Controller) getSellerName(ctx *gin.Context) ([]dtos.Seller, error) {
	queries, err := c.getSearchQuerySeller(ctx)
	if err != nil {
		fmt.Println("bind json error", err)
		utils.ResponseErrorGin(ctx, "bind json error")
		return nil, err
	}

	resp, err := c.SellerService.SearchSeller(queries)
	return resp, nil
}

func (c Controller) NumberOrders(ctx *gin.Context) {
	stepTime := ctx.Query("steptime")
	respnumber := dtos.NumberOrderResponse{}
	respnumber.Steptime = stepTime
	listBranch, _ := c.getBranchnName(ctx)
	listType, _ := c.getTypeName(ctx)
	listSeller, _ := c.getSellerName(ctx)
	t := time.Now()
	t = t.Add(+7 * time.Hour)
	// log.Println(listBranch)
	// log.Println(listType)
	// log.Println(listSeller)
	// getOrderComplatedQuery(ctx *gin.Context, time_s time.Time, typ string, bra string, sel string, steptime string)
	for _, branch := range listBranch {
		queries, err := c.getOrderComplatedQuery(ctx, t, "", branch.Name, "", stepTime)
		if err != nil {
			fmt.Println("bind json error", err)
			utils.ResponseErrorGin(ctx, "bind json error")
			return
		}

		resp, err := c.OrderService.Search(queries)
		if err != nil {
			fmt.Println("search number orders complated error", err)
			utils.ResponseErrorGin(ctx, "search number order complated error")
			return
		}

		data_b := &dtos.NumberOrderInfor{
			Key:   branch.Name,
			Value: int64(len(resp)),
		}

		respnumber.BranchSells = append(respnumber.BranchSells, *data_b)

	}
	for _, typep := range listType {
		queries, err := c.getOrderComplatedQuery(ctx, t, typep.Name, "", "", stepTime)
		if err != nil {
			fmt.Println("bind json error", err)
			utils.ResponseErrorGin(ctx, "bind json error")
			return
		}

		resp, err := c.OrderService.Search(queries)
		if err != nil {
			fmt.Println("search number orders complated error", err)
			utils.ResponseErrorGin(ctx, "search number order complated error")
			return
		}

		data_t := &dtos.NumberOrderInfor{
			Key:   typep.Name,
			Value: int64(len(resp)),
		}
		respnumber.TypeProducts = append(respnumber.TypeProducts, *data_t)

	}
	for _, seller := range listSeller {
		queries, err := c.getOrderComplatedQuery(ctx, t, "", "", seller.Name, stepTime)
		if err != nil {
			fmt.Println("bind json error", err)
			utils.ResponseErrorGin(ctx, "bind json error")
			return
		}

		resp, err := c.OrderService.Search(queries)
		if err != nil {
			fmt.Println("search number orders complated error", err)
			utils.ResponseErrorGin(ctx, "search number order complated error")
			return
		}

		data_s := &dtos.NumberOrderInfor{
			Key:   seller.Name,
			Value: int64(len(resp)),
		}
		respnumber.Sellers = append(respnumber.Sellers, *data_s)

	}

	// List order by day
	if stepTime == "week" {

		for i := 0; i < 7; i++ {
			time := t.AddDate(0, 0, -i)
			queries, err := c.getOrderComplatedQuery(ctx, time, "", "", "", stepTime)
			if err != nil {
				fmt.Println("bind json error", err)
				utils.ResponseErrorGin(ctx, "bind json error")
				return
			}

			resp, err := c.OrderService.Search(queries)
			if err != nil {
				fmt.Println("search number orders complated error", err)
				utils.ResponseErrorGin(ctx, "search number order complated error")
				return
			}
			data_order := &dtos.NumberOrderInfor{
				Key:   time.Format(CommonTimeFormat),
				Value: int64(len(resp)),
			}

			// log.Println(time.Format(CommonTimeFormat))
			// log.Println(len(resp))
			respnumber.Orders = append(respnumber.Orders, *data_order)

		}

		// Branch Sell

	} else if stepTime == "month" {
		for i := 0; i < 31; i++ {
			time := t.AddDate(0, 0, -i)
			queries, err := c.getOrderComplatedQuery(ctx, time, "", "", "", stepTime)
			if err != nil {
				fmt.Println("bind json error", err)
				utils.ResponseErrorGin(ctx, "bind json error")
				return
			}

			resp, err := c.OrderService.Search(queries)
			if err != nil {
				fmt.Println("search number orders complated error", err)
				utils.ResponseErrorGin(ctx, "search number order complated error")
				return
			}
			data_order := &dtos.NumberOrderInfor{
				Key:   time.Format(CommonTimeFormat),
				Value: int64(len(resp)),
			}

			// log.Println(time.Format(CommonTimeFormat))
			// log.Println(len(resp))
			respnumber.Orders = append(respnumber.Orders, *data_order)

		}
	} else if stepTime == "year" {
		for i := 0; i < 12; i++ {
			time := t.AddDate(0, -i, 0)
			// log.Println(time)
			queries, err := c.getOrderComplatedQuery(ctx, time, "", "", "", stepTime)
			if err != nil {
				fmt.Println("bind json error", err)
				utils.ResponseErrorGin(ctx, "bind json error")
				return
			}

			resp, err := c.OrderService.Search(queries)
			if err != nil {
				fmt.Println("search number orders complated error", err)
				utils.ResponseErrorGin(ctx, "search number order complated error")
				return
			}
			data_order := &dtos.NumberOrderInfor{
				Key:   time.Format(CommonTimeFormat),
				Value: int64(len(resp)),
			}

			// log.Println(time.Format(CommonTimeFormat))
			// log.Println(len(resp))
			respnumber.Orders = append(respnumber.Orders, *data_order)

		}
	} else {
		// Do some things
	}
	fmt.Println("search number order complated success")
	log.Println(respnumber)
	utils.ResponseSuccess(ctx, respnumber)
}

// SearchItems
func (c Controller) getSearchItemsQuery(ctx *gin.Context) ([]dtos.SearchItemsQuery, error) {
	result := make([]dtos.SearchItemsQuery, 0)
	ordernumber := ctx.Query("order_number")
	result = append(result, dtos.SearchItemsQuery{
		Key:   "order_number =?",
		Value: ordernumber,
	})
	return result, nil
}

func (c Controller) SearchItems(ctx *gin.Context) {
	queries, err := c.getSearchItemsQuery(ctx)
	if err != nil {
		fmt.Println("bind json error", err)
		utils.ResponseErrorGin(ctx, "bind json error")
		return
	}

	resp, err := c.OrderService.SearchItems(queries)
	if err != nil {
		fmt.Println("search orders error", err)
		utils.ResponseErrorGin(ctx, "search order error")
		return
	}

	fmt.Println("search success")
	utils.ResponseSuccess(ctx, resp)
}

func (c Controller) Search(ctx *gin.Context) {
	queries, err := c.getSearchQuery(ctx)
	if err != nil {
		fmt.Println("bind json error", err)
		utils.ResponseErrorGin(ctx, "bind json error")
		return
	}

	resp, err := c.OrderService.Search(queries)
	if err != nil {
		fmt.Println("search orders error", err)
		utils.ResponseErrorGin(ctx, "search order error")
		return
	}

	fmt.Println("search success")
	utils.ResponseSuccess(ctx, resp)
}

// Scaner Check Order
func (c Controller) Check(ctx *gin.Context) {
	queries, err := c.getSearchQuery(ctx)
	if err != nil {
		fmt.Println("bind json error", err)
		utils.ResponseErrorGin(ctx, "bind json error")
		return
	}
	log.Println(queries)
	resp, err := c.OrderService.Check(queries)
	if err != nil {
		fmt.Println("search orders error", err)
		utils.ResponseErrorGin(ctx, "search order error")
		return
	}

	fmt.Println("search success")
	utils.ResponseSuccess(ctx, resp)
}

// MakeDelay
func (c Controller) MakeDelay(ctx *gin.Context) {
	var request dtos.ChangeStatusToDelay
	err := ctx.ShouldBindJSON(&request)
	if err != nil {
		fmt.Println("bind json error", err)
		utils.ResponseErrorGin(ctx, "bind json error")
		return
	}

	resp, err := c.OrderService.MakeDelay(request.OrderNumber)
	if err != nil {
		fmt.Println("make order delay error", err)
		utils.ResponseErrorGin(ctx, "make order delay error")
		return
	}

	fmt.Println("make order done success")
	utils.ResponseSuccess(ctx, resp)
}

func (c Controller) MakeDone(ctx *gin.Context) {
	var request dtos.ChangeStatusToCompleted
	err := ctx.ShouldBindJSON(&request)
	if err != nil {
		fmt.Println("bind json error", err)
		utils.ResponseErrorGin(ctx, "bind json error")
		return
	}

	resp, err := c.OrderService.MakeCompleted(request.OrderNumber)
	if err != nil {
		fmt.Println("make order done error", err)
		utils.ResponseErrorGin(ctx, "make order done error")
		return
	}

	fmt.Println("make order done success")
	utils.ResponseSuccess(ctx, resp)
}

func (c Controller) Delete(ctx *gin.Context) {
	orderNumber := ctx.Query("order_number")
	if orderNumber == "" {
		utils.ResponseSuccess(ctx, nil)
		return
	}

	err := c.OrderService.Detete(orderNumber)
	if err != nil {
		fmt.Println("delete order error", err)
		utils.ResponseErrorGin(ctx, "delete order error")
		return
	}

	fmt.Println("delete order success")
	utils.ResponseSuccess(ctx, nil)
}

func (c Controller) AddShippingTime(ctx *gin.Context) {
	var request dtos.AddShippingTimeRequest
	err := ctx.ShouldBindJSON(&request)
	if err != nil {
		fmt.Println("bind json error", err)
		utils.ResponseErrorGin(ctx, "bind json error")
		return
	}

	begin, err1 := time.Parse(services.CommonTimeFormat, request.BeginShipping)
	complete, err2 := time.Parse(services.CommonTimeFormat, request.TimeCompleted)
	if err1 != nil || err2 != nil {
		fmt.Println("time parser error", err1, err2)
		utils.ResponseErrorGin(ctx, "time parser error")
		return
	}

	request.BeginShippingReal = &begin
	request.TimeCompletedReal = &complete

	resp, err := c.OrderService.AddShippingTime(request)
	if err != nil {
		fmt.Println("add shipping time error", err)
		utils.ResponseErrorGin(ctx, "add shipping time error")
		return
	}

	fmt.Println("add shipping time success")
	utils.ResponseSuccess(ctx, resp)
}
