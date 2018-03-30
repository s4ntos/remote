// GENERATED CODE - DO NOT EDIT
package main

import (
	"flag"
	"reflect"
	"github.com/revel/revel"
	_ "github.com/mattn/go-sqlite3"
	controllers "github.com/revel/modules/jobs/app/controllers"
	_ "github.com/revel/modules/jobs/app/jobs"
	controllers0 "github.com/revel/modules/static/app/controllers"
	_ "github.com/revel/modules/testrunner/app"
	controllers1 "github.com/revel/modules/testrunner/app/controllers"
	_ "github.com/s4ntos/remote/app"
	controllers2 "github.com/s4ntos/remote/app/controllers"
	_ "github.com/s4ntos/remote/app/jobs"
	models "github.com/s4ntos/remote/app/models"
	tests "github.com/s4ntos/remote/tests"
	"github.com/revel/revel/testing"
)

var (
	runMode    *string = flag.String("runMode", "", "Run mode.")
	port       *int    = flag.Int("port", 0, "By default, read from app.conf")
	importPath *string = flag.String("importPath", "", "Go Import Path for the app.")
	srcPath    *string = flag.String("srcPath", "", "Path to the source root.")

	// So compiler won't complain if the generated code doesn't reference reflect package...
	_ = reflect.Invalid
)

func main() {
	flag.Parse()
	revel.Init(*runMode, *importPath, *srcPath)
	revel.AppLog.Info("Running revel server")
	
	revel.RegisterController((*controllers.Jobs)(nil),
		[]*revel.MethodType{
			&revel.MethodType{
				Name: "Status",
				Args: []*revel.MethodArg{ 
				},
				RenderArgNames: map[int][]string{ 
					73: []string{ 
						"entries",
					},
				},
			},
			
		})
	
	revel.RegisterController((*controllers0.Static)(nil),
		[]*revel.MethodType{
			&revel.MethodType{
				Name: "Serve",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "prefix", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "filepath", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "ServeDir",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "prefix", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "filepath", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "ServeModule",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "moduleName", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "prefix", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "filepath", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "ServeModuleDir",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "moduleName", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "prefix", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "filepath", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			
		})
	
	revel.RegisterController((*controllers1.TestRunner)(nil),
		[]*revel.MethodType{
			&revel.MethodType{
				Name: "Index",
				Args: []*revel.MethodArg{ 
				},
				RenderArgNames: map[int][]string{ 
					76: []string{ 
						"testSuites",
					},
				},
			},
			&revel.MethodType{
				Name: "Suite",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "suite", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "Run",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "suite", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "test", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
					125: []string{ 
					},
				},
			},
			&revel.MethodType{
				Name: "List",
				Args: []*revel.MethodArg{ 
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			
		})
	
	revel.RegisterController((*controllers2.Account)(nil),
		[]*revel.MethodType{
			&revel.MethodType{
				Name: "AddUser",
				Args: []*revel.MethodArg{ 
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "AddAppName",
				Args: []*revel.MethodArg{ 
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "AddRenderMode",
				Args: []*revel.MethodArg{ 
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "Index",
				Args: []*revel.MethodArg{ 
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "Register",
				Args: []*revel.MethodArg{ 
				},
				RenderArgNames: map[int][]string{ 
					104: []string{ 
					},
				},
			},
			&revel.MethodType{
				Name: "SaveUser",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "user", Type: reflect.TypeOf((*models.User)(nil)) },
					&revel.MethodArg{Name: "username", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "name", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "Login",
				Args: []*revel.MethodArg{ 
				},
				RenderArgNames: map[int][]string{ 
					197: []string{ 
						"hasEmailCapability",
					},
				},
			},
			&revel.MethodType{
				Name: "LoginAccount",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "account", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "password", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "remember", Type: reflect.TypeOf((*bool)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "Recover",
				Args: []*revel.MethodArg{ 
				},
				RenderArgNames: map[int][]string{ 
					244: []string{ 
					},
				},
			},
			&revel.MethodType{
				Name: "RetrieveAccount",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "email", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "ConfirmEmail",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "token", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
					303: []string{ 
						"existingProfile",
					},
				},
			},
			&revel.MethodType{
				Name: "PasswordReset",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "token", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "Logout",
				Args: []*revel.MethodArg{ 
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "CheckUserName",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "username", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "currentUsername", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "CheckEmail",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "email", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "currentEmail", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			
		})
	
	revel.RegisterController((*controllers2.Application)(nil),
		[]*revel.MethodType{
			&revel.MethodType{
				Name: "Index",
				Args: []*revel.MethodArg{ 
				},
				RenderArgNames: map[int][]string{ 
					17: []string{ 
					},
				},
			},
			&revel.MethodType{
				Name: "About",
				Args: []*revel.MethodArg{ 
				},
				RenderArgNames: map[int][]string{ 
					21: []string{ 
					},
				},
			},
			&revel.MethodType{
				Name: "Contact",
				Args: []*revel.MethodArg{ 
				},
				RenderArgNames: map[int][]string{ 
					25: []string{ 
					},
				},
			},
			&revel.MethodType{
				Name: "Search",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "query", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "page", Type: reflect.TypeOf((*int)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
					66: []string{ 
						"query",
						"matchedProfiles",
						"page",
						"nextPage",
					},
				},
			},
			&revel.MethodType{
				Name: "SwitchToDesktop",
				Args: []*revel.MethodArg{ 
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "SwitchToMobile",
				Args: []*revel.MethodArg{ 
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			
		})
	
	revel.RegisterController((*controllers2.GorpController)(nil),
		[]*revel.MethodType{
			&revel.MethodType{
				Name: "Begin",
				Args: []*revel.MethodArg{ 
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "Commit",
				Args: []*revel.MethodArg{ 
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "Rollback",
				Args: []*revel.MethodArg{ 
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			
		})
	
	revel.RegisterController((*controllers2.Post)(nil),
		[]*revel.MethodType{
			&revel.MethodType{
				Name: "Show",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "username", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "id", Type: reflect.TypeOf((*int)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
					74: []string{ 
						"title",
						"profile",
						"post",
						"postContentHTML",
						"isOwner",
					},
				},
			},
			&revel.MethodType{
				Name: "Create",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "username", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
					84: []string{ 
						"profile",
					},
				},
			},
			&revel.MethodType{
				Name: "Save",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "username", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "post", Type: reflect.TypeOf((*models.Post)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "Edit",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "username", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "id", Type: reflect.TypeOf((*int)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
					129: []string{ 
						"profile",
						"post",
					},
				},
			},
			&revel.MethodType{
				Name: "Update",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "username", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "id", Type: reflect.TypeOf((*int)(nil)) },
					&revel.MethodArg{Name: "post", Type: reflect.TypeOf((*models.Post)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "Remove",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "username", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "id", Type: reflect.TypeOf((*int)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
					179: []string{ 
						"profile",
						"post",
					},
				},
			},
			&revel.MethodType{
				Name: "Delete",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "username", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "id", Type: reflect.TypeOf((*int)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "Like",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "username", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "id", Type: reflect.TypeOf((*int)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
					215: []string{ 
						"likeResponse",
					},
				},
			},
			
		})
	
	revel.RegisterController((*controllers2.Profile)(nil),
		[]*revel.MethodType{
			&revel.MethodType{
				Name: "Show",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "username", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
					81: []string{ 
						"title",
						"profile",
						"posts",
						"roles",
						"isOwner",
						"isFollowing",
					},
				},
			},
			&revel.MethodType{
				Name: "Settings",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "username", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
					91: []string{ 
						"profile",
					},
				},
			},
			&revel.MethodType{
				Name: "UpdateSettings",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "username", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "profile", Type: reflect.TypeOf((**models.Profile)(nil)) },
					&revel.MethodArg{Name: "verifyPassword", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "Password",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "username", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
					210: []string{ 
					},
				},
			},
			&revel.MethodType{
				Name: "UpdatePassword",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "username", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "password", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "verifyPassword", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "FollowUser",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "username", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
					268: []string{ 
						"followResponse",
					},
				},
			},
			&revel.MethodType{
				Name: "Followers",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "username", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "page", Type: reflect.TypeOf((*int)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
					375: []string{ 
						"profile",
						"isOwner",
						"isFollowing",
						"followerProfiles",
						"page",
						"nextPage",
					},
				},
			},
			&revel.MethodType{
				Name: "Following",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "username", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "page", Type: reflect.TypeOf((*int)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
					405: []string{ 
						"profile",
						"isOwner",
						"isFollowing",
						"followingProfiles",
						"page",
						"nextPage",
					},
				},
			},
			
		})
	
	revel.DefaultValidationKeys = map[string]map[int]string{ 
		"github.com/s4ntos/remote/app/controllers.Account.SaveUser": { 
			118: "user.Password",
			119: "lcPass",
		},
		"github.com/s4ntos/remote/app/controllers.Profile.UpdatePassword": { 
			235: "password",
			236: "verifyPassword",
			237: "verifyPassword",
		},
		"github.com/s4ntos/remote/app/controllers.Profile.UpdateSettings": { 
			114: "User.Password",
			115: "verifyPassword",
			116: "verifyPassword",
		},
		"github.com/s4ntos/remote/app/models.ValidatePostTitle": { 
			29: "title",
			34: "title",
			39: "title",
		},
		"github.com/s4ntos/remote/app/models.ValidateProfileDescription": { 
			98: "description",
		},
		"github.com/s4ntos/remote/app/models.ValidateProfileName": { 
			68: "name",
			73: "name",
			78: "name",
		},
		"github.com/s4ntos/remote/app/models.ValidateProfilePhotoUrl": { 
			104: "photoUrl",
		},
		"github.com/s4ntos/remote/app/models.ValidateProfileSummary": { 
			92: "summary",
		},
		"github.com/s4ntos/remote/app/models.ValidateProfileUserName": { 
			44: "username",
			49: "username",
			54: "username",
		},
		"github.com/s4ntos/remote/app/models.ValidateRoleDescription": { 
			58: "description",
		},
		"github.com/s4ntos/remote/app/models.ValidateRoleName": { 
			39: "role",
			44: "role",
			49: "role",
		},
		"github.com/s4ntos/remote/app/models.ValidateRolePrivileges": { 
			64: "privileges",
		},
		"github.com/s4ntos/remote/app/models.ValidateUserEmail": { 
			41: "email",
			46: "email",
			51: "email",
			56: "email",
		},
		"github.com/s4ntos/remote/app/models.ValidateUserPassword": { 
			62: "password",
			67: "password",
			72: "password",
		},
	}
	testing.TestSuites = []interface{}{ 
		(*tests.AccountsTest)(nil),
		(*tests.ApplicationTest)(nil),
		(*tests.ProfileTest)(nil),
	}

	revel.Run(*port)
}
