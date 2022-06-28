package cdp_test

import (
	"context"
	"fmt"

	"github.com/Unique-AG/rod/lib/cdp"
	"github.com/Unique-AG/rod/lib/launcher"
	"github.com/Unique-AG/rod/lib/proto"
	"github.com/Unique-AG/rod/lib/utils"
	"github.com/ysmood/gson"
)

func ExampleClient() {
	ctx := context.Background()

	// launch a browser
	url := launcher.New().MustLaunch()

	// create a controller
	client := cdp.New(url).MustConnect(ctx)

	go func() {
		for range client.Event() {
			// you must consume the events
		}
	}()

	// Such as call this endpoint on the api doc:
	// https://chromedevtools.github.io/devtools-protocol/tot/Page#method-navigate
	// This will create a new tab and navigate to the test.com
	res, err := client.Call(ctx, "", "Target.createTarget", map[string]string{
		"url": "http://test.com",
	})
	utils.E(err)

	fmt.Println(len(gson.New(res).Get("targetId").Str()))

	// close browser by using the proto lib to encode json
	_ = proto.BrowserClose{}.Call(client)

	// Output: 32
}
