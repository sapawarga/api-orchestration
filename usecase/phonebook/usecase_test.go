package phonebook_test

import (
	"context"
	"fmt"
	"os"
	"reflect"

	mocks "github.com/sapawarga/api-orchestration/mocks/mock_phonebook"
	"github.com/sapawarga/api-orchestration/mocks/mock_phonebook/testcases"
	"github.com/sapawarga/api-orchestration/usecase"
	ucPhonebok "github.com/sapawarga/api-orchestration/usecase/phonebook"

	kitlog "github.com/go-kit/kit/log"
	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"
)

var _ = Describe("Usecase", func() {
	var (
		mockPhonebook *mocks.MockPhonebookI
		usecase       usecase.PhonebookI
	)

	BeforeEach(func() {
		logger := kitlog.NewLogfmtLogger(kitlog.NewSyncWriter(os.Stderr))
		mockSvc := gomock.NewController(GinkgoT())
		mockSvc.Finish()
		mockPhonebook = mocks.NewMockPhonebookI(mockSvc)
		usecase = ucPhonebok.NewUsecase(mockPhonebook, logger)
	})

	// DECLARE UNIT TEST FUNCTION

	var GetListPhonebookLogic = func(idx int) {
		ctx := context.Background()
		data := testcases.GetListPhonebookData[idx]
		mockPhonebook.EXPECT().GetList(ctx, data.GetListRepoParams).Return(data.MockPhonebookRepository.Result, data.MockPhonebookRepository.Error).Times(1)
		resp, err := usecase.GetList(ctx, data.UsecaseParams)
		if err != nil {
			Expect(err).ToNot(BeNil())
			Expect(resp).To(BeNil())
		} else {
			Expect(resp.Metadata.Total).To(Equal(data.MockUsecase.Result.Metadata.Total))
			Expect(err).To(BeNil())
		}
	}

	var GetDetailPhonebookLogic = func(idx int) {
		ctx := context.Background()
		data := testcases.GetDetailPhonebookData[idx]
		mockPhonebook.EXPECT().GetDetail(ctx, data.RepositoryParams).Return(data.MockRepository.Result, data.MockRepository.Error).Times(1)
		resp, err := usecase.GetDetail(ctx, data.UsecaseParams)
		if err != nil {
			Expect(err).ToNot(BeNil())
			Expect(resp).To(BeNil())
		} else {
			Expect(err).To(BeNil())
		}
	}

	var CreatePhonebookLogic = func(idx int) {
		ctx := context.Background()
		data := testcases.CreatePhonebookData[idx]
		mockPhonebook.EXPECT().Insert(ctx, data.RepositoryParam).Return(data.MockRepository).Times(1)
		if err := usecase.Insert(ctx, data.UsecaseParam); err != nil {
			Expect(err).ToNot(BeNil())
		} else {
			Expect(err).To(BeNil())
		}
	}

	var UpdatePhonebookLogic = func(idx int) {
		ctx := context.Background()
		data := testcases.UpdatePhonebookData[idx]
		mockPhonebook.EXPECT().Update(ctx, data.RepositoryParam).Return(data.MockRepository).Times(1)
		if err := usecase.Update(ctx, data.UsecaseParam); err != nil {
			Expect(err).ToNot(BeNil())
		} else {
			Expect(err).To(BeNil())
		}
	}

	var DeletePhonebookLogic = func(idx int) {
		ctx := context.Background()
		data := testcases.DeletePhonebookData[idx]
		mockPhonebook.EXPECT().Delete(ctx, data.RepositoryParam).Return(data.MockRepository).Times(1)
		if err := usecase.Delete(ctx, data.UsecaseParam); err != nil {
			Expect(err).ToNot(BeNil())
		} else {
			Expect(err).To(BeNil())
		}
	}

	var unitTestLogic = map[string]map[string]interface{}{
		"GetList":         {"func": GetListPhonebookLogic, "test_case_count": len(testcases.GetListPhonebookData), "desc": testcases.PhoneBookListDescription()},
		"GetDetail":       {"func": GetDetailPhonebookLogic, "test_case_count": len(testcases.GetDetailPhonebookData), "desc": testcases.PhonebookDetailDescription()},
		"CreatePhonebook": {"func": CreatePhonebookLogic, "test_case_count": len(testcases.CreatePhonebookData), "desc": testcases.InsertPhonebookDescription()},
		"UpdatePhonebook": {"func": UpdatePhonebookLogic, "test_case_count": len(testcases.UpdatePhonebookData), "desc": testcases.UpdatePhonebookDescription()},
		"DeletePhonebook": {"func": DeletePhonebookLogic, "test_case_count": len(testcases.DeletePhonebookData), "desc": testcases.DeletePhonebookDescription()},
	}

	for _, val := range unitTestLogic {
		s := reflect.ValueOf(val["desc"])
		var arr []TableEntry
		for i := 0; i < val["test_case_count"].(int); i++ {
			fmt.Println(s.Index(i).String())
			arr = append(arr, Entry(s.Index(i).String(), i))
		}
		DescribeTable("Function ", val["func"], arr...)
	}
})
