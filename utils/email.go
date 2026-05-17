package utils

import (
	"fmt"
	"os"

	"github.com/resend/resend-go/v2"
)

const welcomeEmailHTML = `
<!DOCTYPE html>
<html lang="en">
<head>
	<meta charset="UTF-8">
	<meta name="viewport" content="width=device-width, initial-scale=1.0">
	<title>Welcome Email</title>
	<style>
		body {
			font-family: Arial, sans-serif;
			color: #333;
			background-color: #f4f4f9;
			padding: 20px;
		}
		.container {
			max-width: 600px;
			margin: auto;
			background-color: #ffffff;
			border-radius: 8px;
			padding: 20px;
			box-shadow: 0px 4px 8px rgba(0, 0, 0, 0.2);
		}
		.header {
			text-align: center;
			color: #4CAF50;
			margin-bottom: 20px;
		}
		.header h1 {
			margin: 0;
		}
		.content p {
			line-height: 1.6;
		}
		.footer {
			text-align: center;
			font-size: 12px;
			color: #888;
			margin-top: 20px;
		}
	</style>
</head>
<body>
	<div class="container">
		<div class="header">
			<h1>Welcome to AlgoLab!</h1>
		</div>
		<div class="content">
			<p>Hello,</p>
			<p>Thank you for signing up for AlgoLab! We're thrilled to have you on board and excited to help you with your coding journey.</p>
			<p>To get started, explore our challenges and start solving problems to improve your skills.</p>
			<p>If you have any questions or need assistance, feel free to reach out to us with this email.</p>
		</div>
		<div class="footer">
			<p>Happy Coding!<br>– The AlgoLab Team</p>
		</div>
	</div>
</body>
</html>
`

func SendEmail(to string) error {
	apiKey := os.Getenv("RESEND_API_KEY")
	if apiKey == "" {
		return fmt.Errorf("RESEND_API_KEY is not set")
	}

	client := resend.NewClient(apiKey)
	params := &resend.SendEmailRequest{
		From:    "onboarding@resend.dev",
		To:      []string{to},
		Subject: "Welcome to AlgoLab!",
		Html:    welcomeEmailHTML,
	}

	_, err := client.Emails.Send(params)
	if err != nil {
		return fmt.Errorf("failed to send welcome email: %w", err)
	}
	return nil
}
