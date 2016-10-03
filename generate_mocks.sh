#!/bin/bash
# generate_mocks - A script to generate the projects mocks.
# It requires golang/mockgen from github.com

rm ./mocks/*_mock.go

$GOPATH/bin/mockgen -source=controllers/account/controller.go -destination=mocks/account_controller_mock.go -package=mocks
$GOPATH/bin/mockgen -source=controllers/budget/controller.go -destination=mocks/budget_controller_mock.go -package=mocks
$GOPATH/bin/mockgen -source=controllers/user/controller.go -destination=mocks/user_controller_mock.go -package=mocks
$GOPATH/bin/mockgen -source=dao/account.go -destination=mocks/account_dao_mock.go -package=mocks
$GOPATH/bin/mockgen -source=dao/budget.go -destination=mocks/budget_dao_mock.go -package=mocks
$GOPATH/bin/mockgen -source=dao/user.go -destination=mocks/user_dao_mock.go -package=mocks