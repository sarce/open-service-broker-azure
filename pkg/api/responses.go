package api

var responseAsyncRequired = []byte(`{ "error": "AsyncRequired", "description": "This service plan requires client support for asynchronous service operations." }`)
var responseServiceIDRequired = []byte(`{ "error": "ServiceIdRequired", "description": "service_id is a required field." }`)
var responsePlanIDRequired = []byte(`{ "error": "PlanIdRequired", "description": "plan_id is a required field." }`)
var responseInvalidServiceID = []byte(`{ "error": "InvalidServiceId", "description": "The provided service_id is invalid." }`)
var responseInvalidPlanID = []byte(`{ "error": "InvalidPlanId", "description": "The provided plan_id is invalid." }`)
var responseEmptyJSON = []byte("{}")