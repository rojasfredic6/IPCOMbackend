package models

import (
	"strings"
)

type JsonResponse struct {
	Organization string      `json:"organization"`
	Users        []UsersData `json:"users"`
}

type UsersData struct {
	Username string   `json:"username"`
	Roles    []string `json:"roles"`
}

func GenerateResJSON(data [][]string) ([]JsonResponse, error) {
	listResp := []JsonResponse{}
	for index, item := range data {
		if index > 0 {
			if len(listResp) > 0 {
				for j, element := range listResp {
					if element.Organization == item[0] {
						for u, user := range element.Users {
							if user.Username == item[1] {
								for _, role := range user.Roles {
									if !strings.Contains(role, item[2]) {
										listResp[j].Users[u].Roles = append(listResp[j].Users[u].Roles, item[2])
									}
								}
							} else {
								var notExist bool
								for _, obj := range listResp[j].Users {
									if obj.Username == item[1] {
										notExist = false
									} else {
										notExist = true
									}
								}

								if notExist {
									newUser := &UsersData{}
									newUser.Username = item[1]
									newUser.Roles = append(newUser.Roles, item[2])
									listResp[j].Users = append(listResp[j].Users, *newUser)
								}
							}
						}
					} else {
						var notExist bool
						for _, obj := range listResp {
							if obj.Organization == item[0] {
								notExist = false
							} else {
								notExist = true
							}
						}
						if notExist {
							newJsonResponse := &JsonResponse{}
							newJsonResponse.Organization = item[0]
							newUser := &UsersData{}
							newUser.Username = item[1]
							newUser.Roles = append(newUser.Roles, item[2])
							newJsonResponse.Users = append(newJsonResponse.Users, *newUser)
							listResp = append(listResp, *newJsonResponse)
						}
					}
				}
			} else {
				newJsonResponse := &JsonResponse{}
				newJsonResponse.Organization = item[0]
				newUser := &UsersData{}
				newUser.Username = item[1]
				newUser.Roles = append(newUser.Roles, item[2])
				newJsonResponse.Users = append(newJsonResponse.Users, *newUser)
				listResp = append(listResp, *newJsonResponse)
			}
		}
	}

	return listResp, nil
}
