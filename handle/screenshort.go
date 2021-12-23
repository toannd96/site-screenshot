package handle

import (
	"bytes"
	"context"
	"io"
	"log"

	emu "github.com/chromedp/cdproto/emulation"
	"github.com/chromedp/cdproto/page"
	cdp "github.com/chromedp/chromedp"
	"github.com/gofiber/fiber/v2"
)

type ScreenShortHandle struct{}

func (s *ScreenShortHandle) ScreenShot(c *fiber.Ctx) error {
	url := c.Query("url")
	ctx, cancel := cdp.NewContext(context.Background())
	defer cancel()

	var buf []byte
	tasks := cdp.Tasks{
		cdp.Navigate(url),
		cdp.ActionFunc(
			func(ctx context.Context) error {
				_, _, _, _, _, contentSize, err := page.GetLayoutMetrics().Do(ctx)
				if err != nil {
					return err
				}
				w, h := contentSize.Width, contentSize.Height
				viewPortFix(ctx, int64(w), int64(h))
				buf, err = page.CaptureScreenshot().
					WithQuality(100).WithClip(&page.Viewport{
					X:      contentSize.X,
					Y:      contentSize.Y,
					Width:  w,
					Height: h,
					Scale:  1,
				}).Do(ctx)
				if err != nil {
					return err
				}
				return nil
			})}

	err := cdp.Run(ctx, tasks)
	if err != nil {
		return err
	}

	c.Response().Header.Set(fiber.HeaderContentDisposition, "inline; filename=screenshot.png")
	c.Response().Header.Set(fiber.HeaderContentType, "image/png")

	_, errCopy := io.Copy(c.Response().BodyWriter(), bytes.NewReader(buf))
	if errCopy != nil {
		return err
	}

	return c.SendStatus(200)
}

func viewPortFix(ctx context.Context, w, h int64) {
	err := emu.SetDeviceMetricsOverride(w, h, 1, false).WithScreenOrientation(
		&emu.ScreenOrientation{
			Type:  emu.OrientationTypePortraitPrimary,
			Angle: 0,
		}).Do(ctx)
	if err != nil {
		log.Fatal(err)
	}
}
