package health

import (
	
	"time"

	"go.temporal.io/sdk/workflow"
)




func Wkflow(ctx workflow.Context)error{

	ao := workflow.ActivityOptions{
        StartToCloseTimeout: time.Minute,
    }

	ctx=workflow.WithActivityOptions(ctx,ao)

	var res string

	err:=workflow.ExecuteActivity(ctx,Healthmeasure).Get(ctx,&res)
	if err!=nil{
		return err
	}

	return nil







}