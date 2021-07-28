package external_svc_test

import (
	"context"
	"github.com/deqode/GoArcc/modules/user/v1/pb"
	"github.com/deqode/GoArcc/util/userinfo"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
	_ "log"
	"time"
)

var _ = Describe("GetUser", func() {
	var (
		usersServer pb.UsersServer
		ctx         context.Context
	)

	// this block will run after each it block
	BeforeEach(func() {
		usersServer = UserServerTest
		ctx = CtxTest
	})

	// negative
	By("internal or external call")
	Context("Get an error when id is wrong", func() {
		It("Return record not found error", func() {
			ui := userinfo.UserInfo{
				ID:          "Id",
				Email:       "Email",
				Sub:         "Sub",
				TokenExpiry: time.Time{},
			}
			newCtx := userinfo.NewContext(context.Background(), ui)
			_, err := usersServer.GetUser(newCtx, nil)
			Expect(err).Should(Equal(gorm.ErrRecordNotFound))
		})
	})
	// positive
	Context("Get a record when id is provided", func() {
		It("should return requested field in the object", func() {
			_, err := usersServer.GetUser(ctx, nil)
			Expect(err).Should(BeNil())
		})
	})
})

var _ = Describe("DeleteUser", func() {
	var (
		userServer pb.UsersServer
		ctx        context.Context
		profile    *pb.User
	)
	BeforeEach(func() {
		profile = Usr
		userServer = UserServerTest
		ctx = CtxTest
	})

	Describe("Deleting a user profile", func() {
		//Negative Cases
		By("By an internal RPC Call")
		Context("Get an error when no user_id provided", func() {
			It("Should return validation error", func() {
				_, err := userServer.DeleteUser(ctx, &pb.DeleteUserRequest{Id: ""})
				Expect(err.(pb.DeleteUserRequestValidationError).Reason()).Should(Equal("value length must be at least 3 runes"))
			})
		})

		Context("Get an error when request is nil", func() {
			It("should return error ", func() {
				_, err := userServer.DeleteUser(ctx, nil)
				Expect(err).Should(Equal(status.Error(codes.FailedPrecondition, "Request can't be nil")))
			})
		})

		Context("Return nil when wrong user_id provided", func() {
			It("should not perform any action in DB and return nil", func() {
				_, err := userServer.DeleteUser(ctx, &pb.DeleteUserRequest{Id: "fejfoeiw0r8290r420"})
				Expect(err).To(BeNil(), "Error")
			})
		})

		//Positive Test Cases
		Context("Return confirmation when record soft deleted in DB", func() {
			It("should return a boolean value", func() {
				_, err := userServer.DeleteUser(ctx, &pb.DeleteUserRequest{Id: profile.Id})
				Expect(err).To(BeNil(), "Error")
			})
		})
	})

})

var _ = Describe("UpdateUser", func() {
	var (
		userServer pb.UsersServer
		//ctx               context.Context
		//profile        *pb.User
	)
	BeforeEach(func() {
		//profile = Usr
		userServer = UserServerTest
		//ctx = CtxTest
	})

	Describe("Update a user profile", func() {
		//Negative Test Cases
		By("By a internal RPC Call")
		Context("Get an error when request object is nil", func() {
			It("should return nil exception", func() {
				_, err := userServer.UpdateUser(context.Background(), &pb.UpdateUserRequest{User: nil})
				Expect(err).Should(Equal(status.Error(codes.FailedPrecondition, "User to update is not provided")))
			})
		})

		Context("Get an error when user profile object is nil", func() {
			It("should return nil exception", func() {
			})
		})

		Context("Get an error when update mask is incorrect ", func() {
			It("should return failed precondition error", func() {
			})
		})

		Context("Get an error when update mask contain id ", func() {
			It("should return failed precondition error", func() {
			})
		})
		Context("Get an error when id is incorrect", func() {
			It("should return failed precondition error", func() {
			})
		})
		Context("Get an error when user profile does not exist", func() {
			It("should return failed precondition error when profile not exist", func() {
			})
		})

		Context("Update user profile when update mask is correct and id is correct", func() {
			It("user profile should be updated successfully", func() {
			})
		})

		Context("Update user profile when update mask is correct and id is correct", func() {
			It("check the response of object that user profile is updated or not", func() {
			})
		})
	})
})
