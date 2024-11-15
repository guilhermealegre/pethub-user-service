-- +migrate Up
-- SQL in section 'Up' is executed when this migration is applied


CREATE TABLE "notification".email_template
(
    id_push_email   SERIAL PRIMARY KEY,
    key   VARCHAR(150) NOT NULL,
    subject         VARCHAR(150) NOT NULL,
    content            TEXT NOT NULL,
    active          BOOLEAN NOT NULL DEFAULT TRUE,
    created_at      timestamptz NOT NULL DEFAULT now(),
    updated_at      timestamptz NOT NULL DEFAULT now()
);


insert into "notification".email_template (key, subject, content) values
(
    'email_signup_confirmation',
    'FitHub email confirmation',
   '<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>FitHub Email Confirmation</title>
</head>
<body>
    <h1>Email Confirmation</h1>
    <p>Thank you for signing up! Please use the code below to confirm your email address and complete the signup process in the app.</p>
    <div style="background-color:#f0f0f0; padding:10px; border-radius: 5px;">
        <p style="font-size: 20px; font-weight: bold;">Confirmation Code:</p>
        <p style="font-size: 24px; color: #007bff;">{{.confirmation_code}}</p>
    </div>
    <p>If you did not sign up for this account, you can ignore this email.</p>
</body>
</html>'
);


-- +migrate Down
-- SQL in section 'Down' is executed when this migration is applied


