package sdek

// https://confluence.cdek.ru/pages/viewpage.action?pageId=29923926

func (c *Client) OrderCreate(orderReq OrderReq) (*OrderRes, error) {

	var err error
	method := "orders?"

	var orderRes *OrderRes
	var orderErr *OrderRes
	_, err = c.post(method, orderReq, &orderRes, &orderErr)
	if err != nil {
		panic(err)
	}

	var errs ErrorsSDK
	if orderErr != nil && len(orderErr.Requests) > 0 {
		for _, request := range orderErr.Requests {
			errs.Errors = append(errs.Errors, request.Errors...)
		}
		if len(errs.Errors) > 0 {
			err = errs
		}
	}
	return orderRes, err

}

func (c *Client) OrderInfoByUUID(uuid string) (*Order, error) {

	method := "orders/" + uuid
	var orderRes *OrderRes
	var orderErr *OrderRes
	_, err := c.get(method, &orderRes, &orderErr)
	if err != nil {
		panic(err)
	}

	return orderRes.Entity, err
}

func (c *Client) OrderInfoByN(uuid string) (*Order, error) {
	method := "orders?cdek_number=" + uuid
	var orderRes *OrderRes
	var orderErr *OrderRes
	_, err := c.get(method, &orderRes, &orderErr)
	if err != nil {
		panic(err)
	}

	return orderRes.Entity, err
}

func (c *Client) OrderInfoByIM(uuid string) (*Order, error) {
	method := "order?_im_number=" + uuid
	var orderRes *OrderRes
	var orderErr *OrderRes
	_, err := c.get(method, &orderRes, &orderErr)
	if err != nil {
		panic(err)
	}

	return orderRes.Entity, err
}
