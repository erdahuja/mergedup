// Package cart contains cart related CRUD functionality.
package cart

import (
	"mergedup/foundation/config"

	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

// Cart manages the set of API's for Cart access.
type Cart struct {
	log   *zap.SugaredLogger
	db    *sqlx.DB
	name  string
	table string
}

// New constructs a record for api access.
func New(dbname, table string,
	log *zap.SugaredLogger,
	db *sqlx.DB,
	mainCfg *config.Configurations,
) (Cart, error) {

	return Cart{
		log:   log,
		db:    db,
		name:  dbname,
		table: table,
	}, nil

}

// // Create adds a KYC record to the database. It returns the created KYC record with
// // fields like ID and DateCreated populated.
// func (k KYC) Create(ctx context.Context, traceID string,
// 	autokycr domain.AutoKYCRequest, kyc *domain.KYC, pid string) (Info, error) {

// 	prd := prepareKYCRecordForDBByDocType(autokycr.UserID,
// 		autokycr.DocumentType, kyc, pid, autokycr.PlatformName.ToString(), autokycr.Medium)
// 	k.log.Printf("%s: %s: %+v", traceID, "kyc.Create", prettyPrint(prd))
// 	return insertIntoDB(ctx, prd, traceID, k.collection, k.name, k.db)
// }

// func prepareKYCRecordForDBByDocType(userID int,
// 	docType string, kyc *domain.KYC, pid string,
// 	platformName string, medium string) Info {
// 	prd := Info{
// 		UserID:       userID,
// 		Medium:       medium,
// 		CreatedAt:    time.Now(),
// 		UpdatedAt:    time.Now(),
// 		ProductID:    pid,
// 		PlatformName: platformName,
// 	}
// 	proofTypes := getProofType(docType)
// 	for _, v := range proofTypes {
// 		if v == domain.UPLOAD_TYPE_ID {
// 			prd.IDProof = kyc
// 		} else if v == domain.UPLOAD_TYPE_ADDRESS {
// 			prd.AddressProof = kyc
// 		} else if v == domain.UPLOAD_TYPE_PAN {
// 			prd.PanProof = kyc
// 		}
// 	}
// 	return prd
// }

// func insertIntoDB(ctx context.Context, prd Info, traceID, collection, dbName string, db database.DB) (Info, error) {
// 	result, err := db.Insert(ctx, dbName, collection, prd, true)
// 	if err != nil || result == "" {
// 		return Info{}, err
// 	}
// 	prd.ID = result
// 	return prd, nil
// }

// // UpdateKYC updates existing KYC record to the database. It returns the updated KYC record
// func (k KYC) UpdateKYC(ctx context.Context, traceID string, autokycr domain.AutoKYCRequest, oldKYC Info, kyc *domain.KYC, pid string) (Info, error) {
// 	proofTypes := getProofType(autokycr.DocumentType)

// 	oldKYC.PlatformName = autokycr.PlatformName.ToString()
// 	oldKYC.Medium = autokycr.Medium
// 	for _, v := range proofTypes {
// 		if v == domain.UPLOAD_TYPE_ID {
// 			oldKYC.IDProof = kyc
// 			oldKYC.IDProof.UpdatedAt = time.Now()
// 		} else if v == domain.UPLOAD_TYPE_ADDRESS {
// 			oldKYC.AddressProof = kyc
// 			oldKYC.AddressProof.UpdatedAt = time.Now()
// 		} else if v == domain.UPLOAD_TYPE_PAN {
// 			oldKYC.PanProof = kyc
// 			oldKYC.PanProof.UpdatedAt = time.Now()
// 		}
// 	}
// 	oldKYC.UpdatedAt = time.Now()
// 	k.log.Printf("%s: %s: %+v", traceID, "kyc.UpdateKYC", prettyPrint(oldKYC))
// 	filter := bson.M{"userID": autokycr.UserID, "productID": pid}
// 	update := bson.M{"$set": oldKYC}
// 	result, err := k.db.Update(ctx, k.name, k.collection, filter, update, true)
// 	if err != nil || !result {
// 		k.log.Printf("%s: %s: Error %+v %s", traceID, "kyc.UpdateKYC", prettyPrint(oldKYC), err)
// 		return Info{}, err
// 	}
// 	return oldKYC, nil
// }

// // Update updates a KYC record in the database with status field filter using docType and userID
// func (k KYC) UpdateStatus(ctx context.Context, traceID, productID string, userID int, req domain.UpdateKYCRequest) error {
// 	k.log.Printf("%s: %s: %d - %+v", traceID, "kyc.Update", userID, prettyPrint(req))
// 	// get kyc
// 	// filter with doc type
// 	// [] --[id, pan] or [pan], or [address]
// 	// update status, msg and modified by
// 	info, err := k.QueryByUserID(ctx, traceID, userID, productID, false)
// 	if err != nil {
// 		return err
// 	}

// 	uploadTypes := findByDocType(info, req.DocType)
// 	if len(uploadTypes) == 0 {
// 		return errors.New(fmt.Sprintf("invalid request, unable to find document: %s in saved kyc record", req.DocType))
// 	}

// 	for _, v := range uploadTypes {
// 		if v == domain.UPLOAD_TYPE_ID {
// 			info.IDProof.Status = req.Status
// 			info.IDProof.StatusMessage = req.StatusMessage
// 			info.IDProof.ModifiedBy = req.ModifiedBy
// 			info.IDProof.UpdatedAt = time.Now()
// 		} else if v == domain.UPLOAD_TYPE_ADDRESS {
// 			info.AddressProof.Status = req.Status
// 			info.AddressProof.StatusMessage = req.StatusMessage
// 			info.AddressProof.ModifiedBy = req.ModifiedBy
// 			info.AddressProof.UpdatedAt = time.Now()
// 		} else if v == domain.UPLOAD_TYPE_PAN {
// 			info.PanProof.Status = req.Status
// 			info.PanProof.StatusMessage = req.StatusMessage
// 			info.PanProof.ModifiedBy = req.ModifiedBy
// 			info.PanProof.UpdatedAt = time.Now()
// 		}
// 	}

// 	objectID, err := primitive.ObjectIDFromHex(info.ID)
// 	if err != nil {
// 		return err
// 	}

// 	filter := bson.M{"_id": objectID}
// 	info.ID = "" // since updating whole existing bson
// 	info.UpdatedAt = time.Now()
// 	update := bson.M{"$set": info}

// 	userProfile, err := k.jwrSDK.GetUserProfile(ctx, userID, 0)
// 	if err != nil {
// 		return errors.Wrap(ErrUserProfie, err.Error())
// 	}

// 	dob, err := userProfile.GetDOB()
// 	if err != nil {
// 		k.log.Printf("TraceId %v Error %v", traceID, err)
// 	}

// 	if req.DocType == domain.DOC_TYPE_AADHAR {

// 		// to make sure panic does not occur because of error out of bound
// 		if info.AddressProof != nil &&
// 			len(info.AddressProof.DocumentID) > 4 {
// 			uniqueKyc := GetUniqueKycObject(InputUniqueKyc{
// 				FirstName:           userProfile.FirstName,
// 				LastName:            userProfile.LastName,
// 				DOB:                 dob,
// 				PinCode:             userProfile.Pin,
// 				LastFourDigitAddhar: info.AddressProof.DocumentID[len(info.AddressProof.DocumentID)-4:],
// 			})

// 			if req.Status == domain.KYC_APPROVED &&
// 				req.DocType == domain.DOC_TYPE_AADHAR {

// 				err = k.uniqueKycAdapter.IncludeUserID(ctx, uniqueKyc, userID)
// 				if err != nil {
// 					k.log.Printf("TraceId %v Error %v", traceID, err)
// 					return ErrRedisDown
// 				}
// 			} else {
// 				err = k.uniqueKycAdapter.Delete(ctx, uniqueKyc, userID)
// 				if err != nil {
// 					k.log.Printf("TraceId %v Error %v", traceID, err)
// 					return ErrRedisDown
// 				}
// 			}
// 		}
// 	}

// 	result, err := k.db.Update(ctx, k.name, k.collection, filter, update, true)
// 	if err != nil || !result {
// 		k.log.Printf("TraceId %v Error %v", traceID, err)
// 		return err
// 	}

// 	k.domainEvents <- domainevents.Model{
// 		UserID:    userID,
// 		ProductID: productID,
// 		RequestID: traceID,
// 		Name:      domainevents.KYC_UPDATE,
// 	}

// 	return nil
// }

// // QueryByUserID gets a KYC record from the database using userID
// func (k KYC) QueryByUserID(ctx context.Context, traceID string, userID int, productID string, usePrimaryDB bool) (Info, error) {

// 	defer nrf.FromContext(ctx).StartSegment("QueryByUserID").End()
// 	var infos []Info
// 	filter := bson.M{"userID": userID, "productID": productID}
// 	err := k.db.Find(ctx, k.name, k.collection, filter, &infos, usePrimaryDB)
// 	if err != nil {
// 		return Info{}, err
// 	}
// 	if len(infos) == 0 {
// 		return Info{}, database.ErrNotFound
// 	}
// 	return infos[0], nil
// }

// func (k KYC) FetchRegionFromPinCode(ctx context.Context, pinCode int) (*domainevents.PinCodeRegion, error) {

// 	if region, exist := k.PinCodeRegion[pinCode]; exist {
// 		return &region, nil
// 	}

// 	var pincodeDbModel []domainevents.PinCodeCollectionModel

// 	filter := bson.M{"Pincode": pinCode}
// 	err := k.db.Find(ctx, k.name, "pincodes", filter, &pincodeDbModel, true)
// 	if err != nil {
// 		return nil, err
// 	}

// 	stateNameDuplicacyCheck := make(map[string]interface{})

// 	for _, eachState := range pincodeDbModel {
// 		stateNameDuplicacyCheck[eachState.StateName] = nil
// 	}

// 	if len(stateNameDuplicacyCheck) > 1 {
// 		return nil, ErrMultipleStateForPincode
// 	}

// 	if len(pincodeDbModel) == 0 {
// 		return nil, database.ErrNotFound
// 	}

// 	var jwrStateName string
// 	var exist bool

// 	if jwrStateName, exist = StateMapFromKyc2JWR[pincodeDbModel[0].StateName]; !exist {
// 		return nil, ErrInvalidStateName
// 	}

// 	output := domainevents.PinCodeRegion{
// 		Pincode: pinCode,
// 		State:   jwrStateName,
// 		City:    pincodeDbModel[0].City,
// 	}

// 	k.PinCodeRegion[pinCode] = output
// 	return &output, nil
// }

// // QueryByDocIDAndApprovedAndSubmitted gets a KYC kyc from the database. It returns the KYC with
// // field matching kyc ID and approved status with approved and submitted status
// func (k KYC) QueryByDocIDAndApprovedAndSubmitted(ctx context.Context, traceID string, docID, productID string) ([]Info, error) {

// 	defer nrf.FromContext(ctx).StartSegment("QueryByDocIDAndApprovedAndSubmitted").End()
// 	var infos []Info
// 	addressProofStatus := bson.A{bson.M{"addressProof.status": domain.KYC_APPROVED}, bson.M{"addressProof.status": domain.KYC_SUBMITTED}}
// 	idProofStatus := bson.A{bson.M{"idProofStatus.status": domain.KYC_APPROVED}, bson.M{"idProofStatus.status": domain.KYC_SUBMITTED}}
// 	panProofStatus := bson.A{bson.M{"panProof.status": domain.KYC_APPROVED}, bson.M{"panProof.status": domain.KYC_SUBMITTED}}
// 	filter := bson.A{bson.M{"addressProof.documentID": docID, "$or": addressProofStatus},
// 		bson.M{"idProof.documentID": docID, "$or": idProofStatus},
// 		bson.M{"panProof.documentID": docID, "$or": panProofStatus}}
// 	err := k.db.Find(ctx, k.name, k.collection, bson.M{"$or": filter, "productID": productID}, &infos, false)
// 	if err != nil {
// 		return infos, err
// 	}
// 	if len(infos) == 0 {
// 		return nil, database.ErrNotFound
// 	}
// 	return infos, nil
// }

// // ISDocIDAlreadyUsed tells if document id is used by any other user in submitted or approved state
// func (k KYC) ISDocIDAlreadyUsed(ctx context.Context, traceID string, documentID, productID, docType string, userID int) (bool, error) {

// 	defer nrf.FromContext(ctx).StartSegment("ISDocIDAlreadyUsed").End()
// 	results, err := k.QueryByDocIDAndApprovedAndSubmitted(ctx, traceID, documentID, productID)
// 	if err != nil {
// 		switch errors.Cause(err) {
// 		case database.ErrNotFound:
// 			break
// 		default:
// 			return false, err
// 		}
// 	}
// 	var isAlreadyUsed bool
// 	for _, v := range results {
// 		if v.UserID != userID {
// 			isAlreadyUsed = true
// 			break
// 		}
// 	}
// 	return isAlreadyUsed, nil
// }

// // ProcessDirectKYC creates a KYC document directly in db
// func (k KYC) ProcessDirectKYC(ctx context.Context,
// 	traceID, productID string,
// 	autokycr domain.AutoKYCRequest,
// 	dmsAndProfileResponse external_service.DMSAndProfileResponse,
// 	ch chan *domain.KYC, _kyc Info,
// 	platformName web.PlatformName) {

// 	defer nrf.FromContext(ctx).StartSegment("ProcessDirectKYC").End()

// 	k.log.Printf("%s:  %s: AutoKYCRequest: %+v\n", traceID, "kyc.ProcessDirectKYC", autokycr)
// 	kb := domain.NewKYCBuilder()
// 	kb.WithBasics(autokycr.DocumentID, autokycr.RecordID, autokycr.DocumentType)

// 	if len(dmsAndProfileResponse.Dms.OCR.Extraction.Identifier) >= 4 &&
// 		autokycr.DocumentType == domain.DOC_TYPE_AADHAR {
// 		firstName, _, lastName := userverification.GetNameParts(dmsAndProfileResponse.Dms.OCR.Extraction.FullName)

// 		uniqueKyc := GetUniqueKycObject(InputUniqueKyc{
// 			FirstName:           firstName,
// 			LastName:            lastName,
// 			DOB:                 dmsAndProfileResponse.Dms.OCR.Extraction.DOB,
// 			PinCode:             dmsAndProfileResponse.Dms.OCR.Extraction.Pin,
// 			LastFourDigitAddhar: dmsAndProfileResponse.Dms.OCR.Extraction.Identifier[len(dmsAndProfileResponse.Dms.OCR.Extraction.Identifier)-4:],
// 		})

// 		duplicateKycUsers, err := k.FindDuplicateKycUsers(ctx, uniqueKyc, traceID)
// 		if err != nil {
// 			k.log.Printf("traceID %v avoid duplicate kyc users error for now %v", traceID, err)
// 		}
// 		kb.WithDuplicateKycUsers(duplicateKycUsers)
// 	}

// 	restriction := domain.Restriction{IsRestricted: false, Message: "Skipped due to direct kyc request by agent"}
// 	var fr domain.FraudResult
// 	fr.Status = domain.DB_VALIDATION_PENDING
// 	fr.Message = "Skipped due to direct kyc request by agent"
// 	kb.WithStatusReason(domain.OCR_INTEGRITY_UNKNOWN, domain.PROFILE_MATCH_NOT_DONE, restriction)
// 	kb.Fraud(fr.Status).WithFraudReason(fr.Message)
// 	kb.WithExtras(restriction, fr, dmsAndProfileResponse.Dms.OCR.Extraction)

// 	kbStatusMessage := "Direct KYC Request By Agent"

// 	if platformName == web.PS_CASHAPP {
// 		kbStatusMessage = "InApp Cash Request"
// 	}

// 	kb.Status(domain.KYC_SUBMITTED).WithMessage(kbStatusMessage)
// 	record := kb.Build()
// 	ch <- record
// }

// // ProcessHZKYC creates a KYC document rom HZ Request if needed
// func (k KYC) ProcessHZKYC(ctx context.Context, traceID, productID string, userID int, req NewHZKYCRequest, profile *profile.UserProfile, record Info) {

// 	defer nrf.FromContext(ctx).StartSegment("ProcessHZKYC").End()
// 	k.log.Printf("%s:  %s: NewHZKYCRequest: %+v\n", traceID, "kyc.ProcessHZKYC", req)
// 	canDoKYC := isKYCCreationAllowed(traceID, k.cfg, &k.log, k.db, k.name, k.collection, req)
// 	if !canDoKYC {
// 		k.log.Printf("%s: cannot do kyc due to state or age restriction from hz request (%v) hence returning\n", traceID, req)
// 		return
// 	}
// 	up := k.BuildResponse(userID, productID, record, profile)
// 	k.log.Printf("UserKYCAndProfileDetails: %+v\n", up)
// 	recordID := fmt.Sprintf("%s%d", "howzat-request-", req.ID)
// 	var jwrPAN, jwrAddress, jwrID *domain.KYC
// 	if req.IDProof != nil && strings.EqualFold(req.IDProof.Status, VERIFIED) {
// 		jwrPAN = getJWRKYCFromHZPAN(&k.log, req.IDProof, up, req.Address, recordID)
// 	}
// 	if req.AddressProof != nil && strings.EqualFold(req.AddressProof.Status, VERIFIED) {
// 		jwrAddress = getJWRKYCFromHZAdress(&k.log, req.AddressProof, up, req.Address, recordID)
// 		jwrID = getJWRKYCFromHZID(&k.log, req.AddressProof, up, req.Address, recordID)
// 	}
// 	updatePAN := jwrPAN != nil
// 	updateAddress := jwrAddress != nil || jwrID != nil
// 	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(time.Second*30))
// 	defer cancel()
// 	if updatePAN {
// 		isAlreadyUsed, err := k.ISDocIDAlreadyUsed(ctx, traceID, jwrPAN.DocumentID, productID, jwrPAN.DocType, userID)
// 		if err != nil || isAlreadyUsed {
// 			k.log.Printf("%s: %s: %+v updatePAN.ISDocIDAlreadyUsed - Error || isAlreadyUsed: %v-%v", traceID, "kyc.ProcessHZKYC", jwrPAN, isAlreadyUsed, err)
// 			return
// 		}
// 	}
// 	if updateAddress {
// 		isAlreadyUsed, err := k.ISDocIDAlreadyUsed(ctx, traceID, jwrAddress.DocumentID, productID, jwrAddress.DocType, userID)
// 		if err != nil || isAlreadyUsed {
// 			k.log.Printf("%s: %s: %+v updateAddress.ISDocIDAlreadyUsed - Error || isAlreadyUsed: %v-%v", traceID, "kyc.ProcessHZKYC", jwrAddress, isAlreadyUsed, err)
// 			return
// 		}
// 	}
// 	if record.IsNew {
// 		info := prepareKYCRecordForDB(userID, jwrPAN, jwrAddress, jwrID, productID)
// 		_, err := insertIntoDB(ctx, info, traceID, k.collection, k.name, k.db)
// 		if err != nil {
// 			k.log.Printf("%s: %s: %+v - Error : %v", traceID, "kyc.ProcessHZKYC.insertIntoDB", info, err)
// 			return
// 		}
// 		k.domainEvents <- domainevents.Model{
// 			UserID:    userID,
// 			ProductID: productID,
// 			Name:      domainevents.KYC_INIT,
// 			RequestID: traceID,
// 		}
// 	} else {
// 		if jwrPAN != nil {
// 			record.PanProof = jwrPAN
// 		}
// 		if jwrAddress != nil {
// 			record.AddressProof = jwrAddress
// 		}
// 		if jwrID != nil {
// 			record.IDProof = jwrID
// 		}
// 		objectID, err := primitive.ObjectIDFromHex(record.ID)
// 		if err != nil {
// 			k.log.Printf("%s: %s: %+v - Error : %v", traceID, "kyc.ProcessHZKYC.ObjectIDFromHex", record, err)
// 			return
// 		}
// 		filter := bson.M{"_id": objectID}
// 		record.ID = "" // since updating whole existing bson
// 		record.UpdatedAt = time.Now()
// 		update := bson.M{"$set": record}
// 		_, err = k.db.Update(ctx, k.name, k.collection, filter, update, true)
// 		if err != nil {
// 			k.log.Printf("%s: %s: %+v - Error : %v", traceID, "kyc.ProcessHZKYC.Update", record, err)
// 			return
// 		}
// 		if updatePAN || updateAddress {
// 			k.domainEvents <- domainevents.Model{
// 				UserID:    userID,
// 				ProductID: productID,
// 				Name:      domainevents.KYC_UPDATE,
// 				RequestID: traceID,
// 			}
// 		}
// 	}
// 	updateProfileFromHZ(traceID, k.profileURI, userID, &k.log, profile, req, updatePAN, updateAddress, k.jwrToken)
// }

// func prepareKYCRecordForDB(userID int, pan, address, id *domain.KYC, pid string) Info {
// 	prd := Info{
// 		UserID:    userID,
// 		CreatedAt: time.Now(),
// 		UpdatedAt: time.Now(),
// 		ProductID: pid,
// 	}
// 	if address != nil {
// 		prd.AddressProof = address
// 	}
// 	if id != nil {
// 		prd.IDProof = id
// 	}
// 	if pan != nil {
// 		prd.PanProof = pan
// 	}
// 	return prd
// }

// // ProcessKYC generates kyc as per input and various verifications and checks
// func (k KYC) ProcessKYC(ctx context.Context,
// 	traceID, productID,
// 	profileURI string,
// 	autokycr domain.AutoKYCRequest,
// 	dmsAndProfileResponse external_service.DMSAndProfileResponse,
// 	_kyc Info,
// ) (*domain.KYC, error) {

// 	defer nrf.FromContext(ctx).StartSegment("ProcessHZKYC").End()
// 	k.log.Printf("%s:  %s: AutoKYCRequest: %+v\n", "kyc.ProcessKYC", traceID, autokycr)
// 	kb := domain.NewKYCBuilder()
// 	kb.WithBasics(autokycr.DocumentID, autokycr.RecordID, autokycr.DocumentType)

// 	if dmsAndProfileResponse.Dms.Status == FAILED {
// 		var errMsg = dmsAndProfileResponse.Dms.ErrMsg
// 		k.log.Printf("Docuemnt Status: %s", dmsAndProfileResponse.Dms.Status)
// 		k.log.Printf("Error Messge: %s", dmsAndProfileResponse.Dms.ErrMsg)

// 		kb.Status(domain.KYC_DECLINED).WithMessage(errMsg)
// 		kb.WithStatusReason(domain.KYC_DECLINED, domain.PROFILE_MATCH_NOT_DONE, domain.Restriction{IsRestricted: false, Message: "Skipped"})
// 		kb.WithExtras(
// 			domain.Restriction{IsRestricted: false, Message: "Skipped"},
// 			domain.FraudResult{
// 				Status:  domain.DB_VALIDATION_UKNOWN,
// 				Message: "Skipped due to: " + errMsg,
// 			}, dmsAndProfileResponse.Dms.OCR.Extraction)
// 		record := kb.Build()

// 		return record, nil
// 	}

// 	if autokycr.PlatformName == web.PS_CASHAPP &&
// 		!strings.EqualFold(dmsAndProfileResponse.Dms.OCR.DocumentType, domain.DOC_TYPE_PAN) {

// 		pincode, err := strconv.Atoi(dmsAndProfileResponse.Dms.OCR.Extraction.Pin)
// 		log.Printf("debug pingcode %v error %v", pincode, err)
// 		if err == nil {
// 			region, err := k.FetchRegionFromPinCode(ctx, pincode)
// 			if err != nil {
// 				log.Printf("Erro PSMG Error while finding state from pincode %v for pin %v", err, pincode)
// 			}

// 			if region != nil &&
// 				autokycr.City == "" {
// 				autokycr.City = region.City
// 			}
// 			if region != nil &&
// 				autokycr.State == "" {
// 				autokycr.State = region.State
// 			}
// 		} else {
// 			log.Printf("ERROR traceID %v unable to load correct pin code %v autkycr %v pincode %v",
// 				traceID, err, autokycr, dmsAndProfileResponse.Dms.OCR.Extraction.Pin)
// 		}
// 	}

// 	if err := userverification.CheckOCRResponseIntegrity(dmsAndProfileResponse.Dms.OCR, autokycr); err != nil {

// 		k.log.Printf("%s:  %s: OCR: %+v Error: %s\n", "kyc.CheckOCRResponseIntegrity",
// 			traceID, dmsAndProfileResponse.Dms.OCR.Extraction, err)

// 		if err.Error() == userverification.ErrOcrInvalidImage.Error() {
// 			kb.Status(domain.KYC_DECLINED).WithMessage(fmt.Sprintf("Invalid Image uploaded by user %s. Kindly check and upload again", autokycr.DocumentType))
// 			kb.WithStatusReason(domain.OCR_INTEGRITY_UNKNOWN, domain.PROFILE_MATCH_NOT_DONE, domain.Restriction{IsRestricted: false, Message: "Skipped"})
// 			kb.WithExtras(
// 				domain.Restriction{IsRestricted: false, Message: "Skipped"},
// 				domain.FraudResult{
// 					Status:  domain.DB_VALIDATION_UKNOWN,
// 					Message: "Skipped due to image uploaded by user was incorrect",
// 				}, dmsAndProfileResponse.Dms.OCR.Extraction)
// 			record := kb.Build()
// 			return record, nil
// 		}

// 		fr := domain.FraudResult{
// 			Status:  domain.DB_VALIDATION_UKNOWN,
// 			Message: "Skipped due to ocr partial and PS_CASHAPP request",
// 		}

// 		if autokycr.PlatformName != web.PS_CASHAPP {
// 			fraudChan := make(chan domain.FraudResult)
// 			// ocr not full, use profile data for kycing
// 			go k.DoFraudCheck(ctx, autokycr, traceID, productID, fraudChan, false)
// 			fr = <-fraudChan
// 			if fr.Status == domain.DB_VALIDATION_FRAUD {
// 				kb.Status(domain.KYC_DECLINED).WithMessage(fmt.Sprintf("We were unable to verify your %s. Kindly check and upload again", autokycr.DocumentType))
// 			} else if fr.Status == domain.DB_VALIDATION_CLIENT_PENDING {
// 				kb.Status(domain.KYC_CLIENT_PENDING).WithMessage("Document failed to be matched due to incorrect input from client")
// 			} else {
// 				kb.Status(domain.KYC_SUBMITTED).WithMessage("Incomplete OCR detected")
// 			}
// 		} else {
// 			kb.Status(domain.KYC_SUBMITTED).WithMessage("PS_CASHAPP request Incomplete OCR detected")
// 		}

// 		kb.Fraud(fr.Status).WithFraudReason(fr.Message)
// 		restriction := domain.Restriction{IsRestricted: false, Message: "Skipped"}
// 		kb.WithStatusReason(domain.OCR_INTEGRITY_PARTIAL, domain.PROFILE_MATCH_NOT_DONE, restriction)
// 		kb.WithExtras(restriction, fr, dmsAndProfileResponse.Dms.OCR.Extraction)
// 		record := kb.Build()
// 		return record, nil
// 	}

// 	firstName, _, lastName := userverification.GetNameParts(dmsAndProfileResponse.Dms.OCR.Extraction.FullName)
// 	var duplicateKycUsers []int
// 	var uniqueKyc UniqueKyc
// 	var err error

// 	if len(dmsAndProfileResponse.Dms.OCR.Extraction.Identifier) > 4 &&
// 		autokycr.DocumentType == domain.DOC_TYPE_AADHAR {
// 		uniqueKyc = GetUniqueKycObject(InputUniqueKyc{
// 			FirstName:           firstName,
// 			LastName:            lastName,
// 			DOB:                 dmsAndProfileResponse.Dms.OCR.Extraction.DOB,
// 			PinCode:             dmsAndProfileResponse.Dms.OCR.Extraction.Pin,
// 			LastFourDigitAddhar: dmsAndProfileResponse.Dms.OCR.Extraction.Identifier[len(dmsAndProfileResponse.Dms.OCR.Extraction.Identifier)-4:],
// 		})

// 		duplicateKycUsers, err = k.FindDuplicateKycUsers(ctx, uniqueKyc, traceID)
// 		if err != nil {
// 			k.log.Printf("traceID %v avoid duplicate kyc users error for now %v", traceID, err)
// 			return nil, ErrRedisDown
// 		}
// 	}

// 	// ocr must be full, go ahead champ :)
// 	shouldReject := make(chan *domain.Restriction)
// 	go checkRestrictions(ctx, &k.log, k.db, k.name, k.collection, traceID, autokycr.DocumentType, dmsAndProfileResponse, k.cfg, shouldReject)
// 	restriction := <-shouldReject
// 	if restriction.Err != nil {
// 		//todo remove fraud check for ps cash app
// 		k.log.Printf("%s:  %s: Error: %s\n", "kyc.CheckRestrictions", traceID, restriction.Err)
// 		// db call must have failed, leave in manual
// 		// do fraud check

// 		fr := domain.FraudResult{
// 			Status:  domain.DB_VALIDATION_UKNOWN,
// 			Message: "Skipped due to restriction test failure and PS_CASHAPP request",
// 		}

// 		if autokycr.PlatformName != web.PS_CASHAPP {
// 			fraudChan := make(chan domain.FraudResult)
// 			go k.DoFraudCheck(ctx, autokycr, traceID, productID, fraudChan, false)
// 			fr := <-fraudChan
// 			if fr.Status == domain.DB_VALIDATION_FRAUD {
// 				kb.Status(domain.KYC_DECLINED).WithMessage(fmt.Sprintf("We were unable to verify your %s. Kindly check and upload again", autokycr.DocumentType))
// 			} else if fr.Status == domain.DB_VALIDATION_CLIENT_PENDING {
// 				kb.Status(domain.KYC_CLIENT_PENDING).WithMessage("Document failed to be matched due to incorrect input from client")
// 			} else {
// 				kb.Status(domain.KYC_SUBMITTED).WithMessage("Unable to do restriction checks")
// 			}
// 		} else {
// 			kb.Status(domain.KYC_SUBMITTED).WithMessage("PS_CASHAPP request")
// 		}

// 		kb.Fraud(fr.Status).WithFraudReason(fr.Message).WithStatusReason(domain.OCR_INTEGRITY_FULL, domain.PROFILE_MATCH_NOT_DONE, domain.Restriction{IsRestricted: false, Message: "Skipped"})
// 		kb.WithExtras(*restriction, fr, dmsAndProfileResponse.Dms.OCR.Extraction)
// 		record := kb.Build()
// 		return record, nil
// 	}
// 	if restriction.IsRestricted {
// 		kb.Status(domain.KYC_DECLINED).WithMessage(restriction.Message)
// 		kb.WithStatusReason(domain.OCR_INTEGRITY_FULL, domain.PROFILE_MATCH_NOT_DONE, *restriction)
// 		kb.Fraud(domain.DB_VALIDATION_PENDING).WithFraudReason("Skipped due to restriction check (user is restricted)")
// 		kb.WithExtras(*restriction, domain.FraudResult{Status: domain.DB_VALIDATION_UKNOWN, Message: "Skipped due to restriction check (user is restricted)"}, dmsAndProfileResponse.Dms.OCR.Extraction)
// 		record := kb.Build()
// 		return record, nil
// 	}

// 	isSecondDoc := false
// 	if autokycr.DocumentType == domain.DOC_TYPE_PAN {
// 		isSecondDoc = isKYCApproved(_kyc.AddressProof) || isKYCApproved(_kyc.IDProof) || isKYCSubmitted(_kyc.AddressProof) || isKYCSubmitted(_kyc.IDProof)
// 	} else {
// 		isSecondDoc = isKYCApproved(_kyc.PanProof) || isKYCSubmitted(_kyc.PanProof)
// 	}

// 	if len(duplicateKycUsers) > 0 {
// 		fraudChan := make(chan domain.FraudResult)
// 		go k.DoFraudCheck(ctx, autokycr, traceID, productID, fraudChan, false)
// 		fr := <-fraudChan

// 		switch status := fr.Status; status {
// 		case domain.DB_VALIDATION_NONFRAUD:
// 			// consider this as duplicate doc
// 			kb.Status(domain.KYC_SUBMITTED).WithMessage(constants.CONFLICT_WITH_EXISTING_USER)
// 			kb.WithStatusReason(domain.OCR_INTEGRITY_FULL, domain.PROFILE_MATCHED, *restriction)
// 		case domain.DB_VALIDATION_FRAUD:
// 			kb.Status(domain.KYC_DECLINED).WithMessage(fmt.Sprintf("We were unable to verify your %s. Kindly check and upload again", autokycr.DocumentType))
// 			kb.WithStatusReason(domain.OCR_INTEGRITY_FULL, domain.PROFILE_MATCHED, *restriction)
// 		case domain.DB_VALIDATION_CLIENT_PENDING:
// 			kb.Status(domain.KYC_CLIENT_PENDING).WithMessage("Document failed to be matched due to incorrect input from client")
// 			kb.WithStatusReason(domain.OCR_INTEGRITY_FULL, domain.PROFILE_MATCHED, *restriction)
// 		case domain.DB_VALIDATION_PENDING:
// 			// consider this as duplicate doc
// 			kb.Status(domain.KYC_SUBMITTED).WithMessage(constants.CONFLICT_WITH_EXISTING_USER)
// 			kb.WithStatusReason(domain.OCR_INTEGRITY_FULL, domain.PROFILE_MATCHED, *restriction)
// 		case domain.DB_VALIDATION_UKNOWN:
// 			kb.Status(domain.KYC_SUBMITTED).WithMessage("The user status is UNKNOWN (slow response) on govt database but is ocr matched.")
// 			kb.WithStatusReason(domain.OCR_INTEGRITY_FULL, domain.PROFILE_MATCHED, *restriction)
// 		}

// 		kb.Fraud(fr.Status).WithFraudReason(fr.Message)
// 		kb.WithExtras(*restriction, fr, dmsAndProfileResponse.Dms.OCR.Extraction)
// 		kb.WithDuplicateKycUsers(duplicateKycUsers)
// 		record := kb.Build()
// 		return record, nil
// 	}

// 	mm := userverification.GetMismatches(autokycr, dmsAndProfileResponse.Dms.OCR, isSecondDoc, *k.cfg, &k.log)
// 	isProfileMismatched := len(mm) > 0
// 	k.log.Printf("%s: %s: Mismmatches: %v isSecondDoc:%v OCR: %+v AutoKYC: %v\n", "kyc.isProfileMismatched", traceID, mm, isSecondDoc, dmsAndProfileResponse.Dms.OCR, autokycr)
// 	if !isProfileMismatched {
// 		fraudChan := make(chan domain.FraudResult)
// 		go k.DoFraudCheck(ctx, autokycr, traceID, productID, fraudChan, false)
// 		fr := <-fraudChan
// 		switch status := fr.Status; status {
// 		case domain.DB_VALIDATION_NONFRAUD:
// 			kb.Status(domain.KYC_APPROVED).
// 				WithMessage("Document proof matched with official government document.")

// 			kb.WithStatusReason(domain.OCR_INTEGRITY_FULL, domain.PROFILE_MATCHED, *restriction)

// 			if autokycr.DocumentType == domain.DOC_TYPE_AADHAR {
// 				err = k.uniqueKycAdapter.IncludeUserID(ctx, uniqueKyc, autokycr.UserID)
// 				if err != nil {
// 					k.log.Printf("TraceId %v Error %v", traceID, err)
// 					return nil, ErrRedisDown
// 				}
// 			}

// 		case domain.DB_VALIDATION_FRAUD:
// 			kb.Status(domain.KYC_DECLINED).WithMessage(fmt.Sprintf("We were unable to verify your %s. Kindly check and upload again", autokycr.DocumentType))
// 			kb.WithStatusReason(domain.OCR_INTEGRITY_FULL, domain.PROFILE_MATCHED, *restriction)
// 		case domain.DB_VALIDATION_CLIENT_PENDING:
// 			kb.Status(domain.KYC_CLIENT_PENDING).WithMessage("Document failed to be matched due to incorrect input from client")
// 			kb.WithStatusReason(domain.OCR_INTEGRITY_FULL, domain.PROFILE_MATCHED, *restriction)
// 		case domain.DB_VALIDATION_PENDING:
// 			kb.Status(domain.KYC_APPROVED).WithMessage("The user status is PENDING on govt database but is ocr matched.")
// 			kb.WithStatusReason(domain.OCR_INTEGRITY_FULL, domain.PROFILE_MATCHED, *restriction)

// 			if autokycr.DocumentType == domain.DOC_TYPE_AADHAR {
// 				err = k.uniqueKycAdapter.IncludeUserID(ctx, uniqueKyc, autokycr.UserID)
// 				if err != nil {
// 					k.log.Printf("TraceId %v Error %v", traceID, err)
// 					kb.WithMessage("redis down duplicate kyc users check turn off" + err.Error())
// 				}
// 			}

// 		case domain.DB_VALIDATION_UKNOWN:
// 			kb.Status(domain.KYC_SUBMITTED).WithMessage("The user status is UNKNOWN (slow response) on govt database but is ocr matched.")
// 			kb.WithStatusReason(domain.OCR_INTEGRITY_FULL, domain.PROFILE_MATCHED, *restriction)
// 		}
// 		kb.Fraud(fr.Status).WithFraudReason(fr.Message)
// 		kb.WithExtras(*restriction, fr, dmsAndProfileResponse.Dms.OCR.Extraction)
// 		record := kb.Build()
// 		return record, nil
// 	}
// 	// update profile
// 	// override autokycr
// 	// var prof profile.UserProfile
// 	for _, v := range mm {
// 		switch name := v.Name; name {
// 		case "Pin":
// 			dmsAndProfileResponse.Profile.Pin = v.NewValue
// 			autokycr.Pin = v.NewValue
// 		case "DOB":
// 			dmsAndProfileResponse.Profile.DOB = userverification.GetDOBInMs(v.NewValue)
// 			autokycr.DOB = v.NewValue
// 		case "FullName":
// 			dmsAndProfileResponse.Profile.FirstName, dmsAndProfileResponse.Profile.MiddleName, dmsAndProfileResponse.Profile.LastName = userverification.GetNameParts(v.NewValue)
// 			autokycr.FullName = v.NewValue
// 		case "documentID":
// 			autokycr.DocumentID = v.NewValue
// 			kb.DocumentID(v.NewValue)
// 		case "City":
// 			dmsAndProfileResponse.Profile.City = v.NewValue
// 		case "State":
// 			dmsAndProfileResponse.Profile.State = v.NewValue
// 		case "Address":
// 			dmsAndProfileResponse.Profile.Address = v.NewValue
// 		}
// 	}
// 	kb.WithBasics(autokycr.DocumentID, autokycr.RecordID, autokycr.DocumentType)
// 	fraudChan := make(chan domain.FraudResult)
// 	go k.DoFraudCheck(ctx, autokycr, traceID, productID, fraudChan, true)
// 	fr := <-fraudChan
// 	if fr.Status == domain.DB_VALIDATION_FRAUD || fr.Status == domain.DB_VALIDATION_CLIENT_PENDING {
// 		if fr.Status == domain.DB_VALIDATION_FRAUD {
// 			kb.Status(domain.KYC_DECLINED).WithMessage(fmt.Sprintf("We were unable to verify your %s. Kindly check and upload again", autokycr.DocumentType))
// 		} else if fr.Status == domain.DB_VALIDATION_CLIENT_PENDING {
// 			kb.Status(domain.KYC_CLIENT_PENDING).WithMessage("Document failed to be matched due to incorrect input from client")
// 		}
// 		kb.WithStatusReason(domain.OCR_INTEGRITY_FULL, domain.PROFILE_UNMATCHED, *restriction)
// 		kb.Fraud(fr.Status).WithFraudReason(fr.Message)
// 		kb.WithMismatchedFields(mm)
// 		kb.WithExtras(*restriction, fr, dmsAndProfileResponse.Dms.OCR.Extraction)
// 		record := kb.Build()
// 		return record, nil
// 	}
// 	// if doc id is mm, update in kyc entry ie "nd" and "autokycr"
// 	// dont update if kyc doc is second document
// 	// check if this user id already has document uploaded in submitted or approved status
// 	if isSecondDoc {
// 		if fr.Status == domain.DB_VALIDATION_NONFRAUD || fr.Status == domain.DB_VALIDATION_PENDING {
// 			// prevs doc is approved then reject else submitted
// 			s := domain.KYC_SUBMITTED
// 			m := "profile mismatch with existing submitted kyc of userID " + strconv.Itoa(_kyc.UserID)
// 			if (autokycr.DocumentType == domain.DOC_TYPE_PAN &&
// 				isKYCApproved(_kyc.AddressProof) ||
// 				isKYCApproved(_kyc.IDProof)) ||
// 				isKYCApproved(_kyc.PanProof) {
// 				s = domain.KYC_DECLINED
// 				m = "Rejecting due to profile mismatch with existing approved document"
// 			}
// 			kb.Status(s).WithMessage(m)
// 		} else if fr.Status == domain.DB_VALIDATION_FRAUD {
// 			kb.Status(domain.KYC_DECLINED).WithMessage(fmt.Sprintf("We were unable to verify your %s. Kindly check and upload again", autokycr.DocumentType))
// 		} else if fr.Status == domain.DB_VALIDATION_CLIENT_PENDING {
// 			kb.Status(domain.KYC_CLIENT_PENDING).WithMessage("Document failed to be matched due to incorrect input from client")
// 		} else {
// 			kb.Status(domain.KYC_SUBMITTED)
// 		}
// 		kb.WithStatusReason(domain.OCR_INTEGRITY_FULL, domain.PROFILE_UNMATCHED, *restriction)
// 	} else {
// 		if fr.Status == domain.DB_VALIDATION_NONFRAUD || fr.Status == domain.DB_VALIDATION_PENDING {
// 			UpdateProfile(traceID, profileURI, autokycr.UserID, dmsAndProfileResponse.Profile, &k.log, k.jwrToken)
// 			kb.WithStatusReason(domain.OCR_INTEGRITY_FULL, domain.PROFILE_AND_OVERRIDE, *restriction)
// 			kb.Status(domain.KYC_APPROVED)

// 			if autokycr.DocumentType == domain.DOC_TYPE_AADHAR {
// 				err = k.uniqueKycAdapter.IncludeUserID(ctx, uniqueKyc, autokycr.UserID)
// 				if err != nil {
// 					k.log.Printf("TraceId %v Error %v", traceID, err)
// 					return nil, ErrRedisDown
// 				}
// 			}

// 		} else if fr.Status == domain.DB_VALIDATION_FRAUD {
// 			kb.Status(domain.KYC_DECLINED).WithMessage(fmt.Sprintf("We were unable to verify your %s. Kindly check and upload again", autokycr.DocumentType))
// 			kb.WithStatusReason(domain.OCR_INTEGRITY_FULL, domain.PROFILE_UNMATCHED, *restriction)
// 		} else if fr.Status == domain.DB_VALIDATION_CLIENT_PENDING {
// 			kb.Status(domain.KYC_CLIENT_PENDING).WithMessage("Document failed to be matched due to incorrect input from client")
// 			kb.WithStatusReason(domain.OCR_INTEGRITY_FULL, domain.PROFILE_UNMATCHED, *restriction)
// 		} else {
// 			kb.Status(domain.KYC_SUBMITTED)
// 			kb.WithStatusReason(domain.OCR_INTEGRITY_FULL, domain.PROFILE_UNMATCHED, *restriction)
// 		}
// 	}
// 	kb.Fraud(fr.Status).WithFraudReason(fr.Message)
// 	kb.WithMismatchedFields(mm)
// 	kb.WithExtras(*restriction, fr, dmsAndProfileResponse.Dms.OCR.Extraction)
// 	record := kb.Build()
// 	return record, nil
// }

// func (k KYC) FindPendingKycDoc(ctx context.Context, from, to time.Time, limit, page int) ([]Info, error) {

// 	users := make([]Info, 0)
// 	findOptions := options.Find()
// 	findOptions.SetLimit(int64(limit))
// 	skip := int64(page * limit)
// 	findOptions.SetSkip(skip)

// 	// add unique index in mongodb for userID and productID
// 	err := k.db.FindWithOpts(ctx, k.name, "kyc",
// 		bson.M{"addressProof.status": domain.KYC_CLIENT_PENDING,
// 			"addressProof.updatedAt": bson.M{"$gte": from, "$lte": to}}, &users, findOptions, true)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return users, nil
// }

// func (k KYC) FetchKycDoc(ctx context.Context, from, to time.Time, limit, page int) ([]Info, error) {

// 	users := make([]Info, 0)
// 	findOptions := options.Find()
// 	findOptions.SetLimit(int64(limit))
// 	skip := int64(page * limit)
// 	findOptions.SetSkip(skip)

// 	// add unique index in mongodb for userID and productID
// 	err := k.db.FindWithOpts(ctx, k.name, "kyc",
// 		bson.M{"createdAt": bson.M{"$gte": from, "$lte": to}}, &users, findOptions, true)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return users, nil
// }

// func (k KYC) DoFraudCheck(ctx context.Context, autokycr domain.AutoKYCRequest, traceID, productId string, out chan domain.FraudResult, withUpdateProfile bool) {

// 	defer nrf.FromContext(ctx).StartSegment("DoFraudCheck").End()
// 	if !k.cfg.FraudEnabled {
// 		out <- domain.FraudResult{
// 			Status:  domain.DB_VALIDATION_PENDING,
// 			Message: "DB validation is turned off",
// 		}
// 	}
// 	if supported := IsFraudCheckDocTypeValid(autokycr.DocumentType, k.cfg); !supported {
// 		out <- domain.FraudResult{
// 			Status:  domain.DB_VALIDATION_PENDING,
// 			Message: "DB validation not supported for current document",
// 		}
// 	}
// 	if skipped := isCaptchaAvailable(autokycr.Captcha, autokycr.SessionID); !skipped {
// 		out <- domain.FraudResult{
// 			Status:  domain.DB_VALIDATION_PENDING,
// 			Message: "DB validation skipped for current document due to captcha unavailable",
// 		}
// 	}
// 	fraudChannel := make(chan domain.FraudResult)
// 	go fraudcheck.FraudCheck(ctx, autokycr, *k.cfg, fraudChannel, &k.log, k.rc, k.aadharTimeout)
// 	resultSent := false
// 	done := false
// 	for !done {
// 		select {
// 		case fs := <-fraudChannel:
// 			if !resultSent {
// 				log.Printf("Received db validation result in time for traceID %s, docID %s and userID %d with Status %v \n", traceID, autokycr.DocumentID, autokycr.UserID, prettyPrint(fs))
// 				out <- fs
// 				resultSent = true
// 				close(out)
// 			} else {
// 				log.Printf("Received db validation result late for traceID %s, docID %s and userID %d with Status %v \n", traceID, autokycr.DocumentID, autokycr.UserID, prettyPrint(fs))
// 				k.UpdateFraudStatusWithDocType(ctx, traceID, productId, autokycr.UserID, autokycr.DocumentType, fs, false, withUpdateProfile)
// 			}
// 			done = true
// 		case <-time.After(10 * time.Second):
// 			if !resultSent {
// 				fr := domain.FraudResult{
// 					Status:  domain.DB_VALIDATION_UKNOWN,
// 					Message: "DB validation taking too long",
// 				}
// 				out <- fr
// 				close(out)
// 				resultSent = true
// 				ctx = context.Background()
// 			}
// 		}
// 	}
// }

// // UpdateFraudStatusWithDocType updates a KYC record in the database with status field filter using userID and doc type
// func (k KYC) UpdateFraudStatusWithDocType(ctx context.Context, traceID, productID string, userID int, docType string, fr domain.FraudResult, isEarly, withUpdateProfile bool) error {

// 	defer nrf.FromContext(ctx).StartSegment("UpdateFraudStatusWithDocType").End()
// 	filter := bson.M{"userID": userID, "productID": productID}
// 	var update bson.M

// 	uploadtypes := getProofType(docType)
// 	var rec Info
// 	var err error
// 	isSecondDoc := false
// 	if fr.Status == domain.DB_VALIDATION_NONFRAUD || fr.Status == domain.DB_VALIDATION_PENDING || fr.Status == domain.DB_VALIDATION_CLIENT_PENDING {
// 		rec, err = k.QueryByUserID(ctx, traceID, userID, productID, true)
// 		if err != nil {
// 			if err == database.ErrNotFound {
// 				isSecondDoc = false
// 			} else {
// 				k.log.Printf("Error %s : UpdateFraudStatusWithDocType : finding kyc for user id %v %v\n", traceID, userID, err)
// 				return err
// 			}
// 		} else {
// 			if docType == domain.DOC_TYPE_PAN {
// 				isSecondDoc = isKYCApproved(rec.AddressProof) || isKYCApproved(rec.IDProof) || isKYCSubmitted(rec.AddressProof) || isKYCSubmitted(rec.IDProof)
// 			} else {
// 				isSecondDoc = isKYCApproved(rec.PanProof) || isKYCSubmitted(rec.PanProof)
// 			}
// 		}
// 	}
// 	var duplicateKycUsers []int
// 	var uniqueKyc UniqueKyc

// 	if docType == domain.DOC_TYPE_AADHAR {

// 		userProfile, err := getCorrectProfileInfo(&k.log, rec.GetMismatches(domain.UPLOAD_TYPE_ADDRESS), k.profileURI, traceID, userID, k.jwrToken)
// 		if err != nil {
// 			// can't return error from here as in psrmg case there might be chance we didn't got any info from jwr
// 			k.log.Println(err)
// 		} else {
// 			// smelly code
// 			if len(rec.AddressProof.DocumentID) > 4 {
// 				dob, err := changeEpochToString(userProfile.DOB)
// 				if err != nil {
// 					k.log.Printf("TraceId %v Error %v", traceID, err)
// 					return err
// 				}

// 				uniqueKyc = GetUniqueKycObject(InputUniqueKyc{
// 					FirstName:           userProfile.FirstName,
// 					LastName:            userProfile.LastName,
// 					DOB:                 dob,
// 					PinCode:             userProfile.Pin,
// 					LastFourDigitAddhar: rec.AddressProof.DocumentID[len(rec.AddressProof.DocumentID)-4:],
// 				})

// 				duplicateKycUsers, err = k.FindDuplicateKycUsers(ctx, uniqueKyc, traceID)
// 				if err != nil {
// 					k.log.Printf("traceID %v avoid duplicate kyc users error for now %v", traceID, err)
// 					return ErrRedisDown
// 				}
// 			}
// 		}
// 	}

// 	profileUpdated := false
// 	isDbUpdated := false
// 	for _, u := range uploadtypes {
// 		kycStatus := domain.KYC_SUBMITTED
// 		kycStatusMessage := getKYCStatusMessageFromFraud(fr, docType)
// 		frStatus := fr.Status
// 		frMessage := fr.Message
// 		if !isEarly {
// 			frMessage += " | After slow db response"
// 		}
// 		statusKey := u + "Proof.dbValidation"
// 		statusMsg := u + "Proof.dbValidationReason"
// 		statusReason := u + "Proof.statusReason"
// 		updatedAtKey := u + "Proof.updatedAt"
// 		updatedAtValue := time.Now()
// 		kycStatusKey := u + "Proof.status"
// 		kycStatusMessageKey := u + "Proof.statusMessage"
// 		kycDuplicateKysUsersMessageKey := u + "Proof.duplicateKycUsers"

// 		var oldKYCProfileStatus string
// 		switch fr.Status {
// 		case domain.DB_VALIDATION_FRAUD:
// 			kycStatus = domain.KYC_DECLINED
// 			kycStatusMessage = getKYCStatusMessageFromFraud(fr, docType)
// 			update = bson.M{statusKey: frStatus,
// 				statusMsg:           frMessage,
// 				kycStatusKey:        kycStatus,
// 				kycStatusMessageKey: kycStatusMessage,
// 				updatedAtKey:        updatedAtValue,
// 			}
// 		case domain.DB_VALIDATION_UKNOWN:
// 			kycStatusMessage = getKYCStatusMessageFromFraud(fr, docType)
// 			kycStatus = domain.KYC_SUBMITTED
// 			if len(duplicateKycUsers) > 0 {
// 				kycStatusMessage = constants.CONFLICT_WITH_EXISTING_USER
// 			}
// 			update = bson.M{statusKey: frStatus,
// 				statusMsg:                      frMessage,
// 				kycStatusKey:                   kycStatus,
// 				kycStatusMessageKey:            kycStatusMessage,
// 				updatedAtKey:                   updatedAtValue,
// 				kycDuplicateKysUsersMessageKey: duplicateKycUsers,
// 			}
// 		case domain.DB_VALIDATION_NONFRAUD, domain.DB_VALIDATION_PENDING, domain.DB_VALIDATION_CLIENT_PENDING:
// 			if isEarly && frStatus == domain.DB_VALIDATION_CLIENT_PENDING {
// 				kycStatus = domain.KYC_CLIENT_PENDING
// 				update = bson.M{statusKey: frStatus,
// 					statusMsg:    frMessage,
// 					updatedAtKey: updatedAtValue,
// 				}
// 			} else {
// 				_kyc := rec.GetKYC(u)
// 				if _kyc == nil {
// 					continue
// 				}
// 				oldKYCReason := _kyc.StatusReason
// 				if _kyc.Status == domain.KYC_DECLINED {
// 					// no kyc status update
// 					update = bson.M{statusKey: frStatus,
// 						statusMsg:           frMessage,
// 						kycStatusKey:        domain.KYC_DECLINED,
// 						kycStatusMessageKey: kycStatusMessage,
// 						updatedAtKey:        updatedAtValue,
// 					}
// 				} else {
// 					shouldApprove := !oldKYCReason.RestrictionStatus && oldKYCReason.OCRIntegrity == "FULL"

// 					if shouldApprove && len(duplicateKycUsers) == 0 {
// 						kycStatus = domain.KYC_APPROVED
// 						if withUpdateProfile {
// 							if oldKYCReason.ProfileStatus == domain.PROFILE_UNMATCHED {
// 								oldKYCProfileStatus = domain.PROFILE_UNMATCHED
// 								oldKYCReason.ProfileStatus = domain.PROFILE_AND_OVERRIDE
// 							}
// 						}
// 						if uniqueKyc != "" && docType == domain.DOC_TYPE_AADHAR {
// 							err = k.uniqueKycAdapter.IncludeUserID(ctx, uniqueKyc, userID)
// 							if err != nil {
// 								k.log.Printf("TraceId %v Error %v", traceID, err)
// 								return err
// 							}
// 						}

// 					} else {
// 						kycStatus = domain.KYC_SUBMITTED
// 					}
// 					if len(duplicateKycUsers) > 0 {
// 						kycStatusMessage = constants.CONFLICT_WITH_EXISTING_USER
// 					}
// 					update = bson.M{statusKey: frStatus,
// 						statusMsg:                      frMessage,
// 						statusReason:                   oldKYCReason,
// 						kycStatusKey:                   kycStatus,
// 						kycStatusMessageKey:            kycStatusMessage,
// 						updatedAtKey:                   updatedAtValue,
// 						kycDuplicateKysUsersMessageKey: duplicateKycUsers,
// 					}
// 				}
// 			}
// 		}
// 		log.Printf("%s: %s: %d - result: %s - update: %s\n", traceID, "kyc.UpdateFraudStatusWithDocType", userID, prettyPrint(fr), prettyPrint(update))
// 		_, err := k.db.Update(ctx, k.name, k.collection, filter, bson.M{"$set": update}, true)
// 		if err != nil {
// 			return errors.Wrapf(err, "KYC Update request filter: %+v, update: %+v", filter, fr)
// 		}
// 		isDbUpdated = true
// 		kycPart := rec.GetKYC(u)
// 		var sr string
// 		if kycPart != nil {
// 			if kycPart.StatusReason != nil {
// 				sr = kycPart.StatusReason.ProfileStatus
// 			}
// 		}

// 		log.Printf("%s: %s: %d - withUpdateProfile: %v, isSecondDoc: %v, kycStatus: %s, ProfileStatus: %v, profileUpdated: %v\n", traceID, "kyc.UpdateFraudStatusWithDocType", userID, withUpdateProfile, isSecondDoc, kycStatus, sr, profileUpdated)
// 		if withUpdateProfile {
// 			// send profile update
// 			if !isSecondDoc && kycStatus == domain.KYC_APPROVED && !profileUpdated && strings.EqualFold(oldKYCProfileStatus, domain.PROFILE_UNMATCHED) {
// 				go getAndUpdateProfile(&k.log, k.profileURI, traceID, userID, rec.GetMismatches(u), k.jwrToken)
// 				profileUpdated = true
// 			}
// 		}
// 	}
// 	if isDbUpdated {
// 		k.domainEvents <- domainevents.Model{
// 			UserID:    userID,
// 			ProductID: productID,
// 			Name:      domainevents.KYC_UPDATE,
// 			RequestID: traceID,
// 		}
// 	}
// 	return nil
// }

// func changeEpochToString(u int64) (string, error) {
// 	t := time.Unix(0, int64(u)*int64(time.Millisecond))

// 	loc, err := time.LoadLocation("Asia/Kolkata")
// 	if err != nil {
// 		return "", err
// 	}

// 	t = t.In(loc)
// 	s := t.Format("02/01/2006")
// 	return s, nil

// }

// func (k KYC) FetchDMSAndProfile(ctx context.Context, traceID, dmsuri, profileuri, recordID string, userID int) external_service.DMSAndProfileResponse {

// 	defer nrf.FromContext(ctx).StartSegment("FetchDMSAndProfile").End()
// 	requests := []external_service.Request{{URL: dmsuri, To: "dms", ID: recordID}, {URL: profileuri, To: "profile", ID: strconv.Itoa(userID)}}
// 	var dmsResponse domain.DMS
// 	var errDMS error
// 	var profileResponse profile.UserProfile
// 	var errProfile error
// 	wg := sync.WaitGroup{}
// 	for _, req := range requests {
// 		wg.Add(1)
// 		go func(req external_service.Request) {
// 			if req.To == "dms" {
// 				dmsResponse, errDMS = external_service.GETForDMS(req, &k.log)
// 				k.log.Printf("%s:  %s: response: %+v, error: %+v", "kyc.GETForDMS", traceID, dmsResponse, errDMS)
// 			} else if req.To == "profile" {
// 				profileResponse, errProfile = external_service.GETForProfile(req, &k.log, k.jwrToken)
// 				k.log.Printf("%s:  %s: response: %+v, error: %+v", "kyc.GETForProfile", traceID, profileResponse, errProfile)
// 			}
// 			wg.Done()
// 		}(req)
// 	}
// 	wg.Wait()
// 	return external_service.DMSAndProfileResponse{
// 		ErrProfile: errProfile,
// 		ErrDMS:     errDMS,
// 		Profile:    profileResponse,
// 		Dms:        dmsResponse,
// 	}
// }

// func (k KYC) FetchProfile(ctx context.Context, traceID, profileuri string, userID int) (*profile.UserProfile, error) {

// 	defer nrf.FromContext(ctx).StartSegment("FetchProfile").End()
// 	request := external_service.Request{URL: profileuri, To: "profile", ID: strconv.Itoa(userID)}
// 	var profileResponse profile.UserProfile
// 	var errProfile error
// 	profileResponse, errProfile = external_service.GETForProfile(request, &k.log, k.jwrToken)
// 	k.log.Printf("%s:  %s: response: %+v, error: %+v", "kyc.GETForProfile", traceID, profileResponse, errProfile)
// 	return &profileResponse, errProfile
// }

// func (k KYC) BuildResponse(userID int, productID string, record Info, _profile *profile.UserProfile) ProfileAndKYC {
// 	var profileAndKYC ProfileAndKYC
// 	profileAndKYC.UserID = userID
// 	profileAndKYC.ProductID = productID
// 	profileAndKYC.Profile = _profile
// 	profileAndKYC.isNewKYC = record.IsNew
// 	if record.AddressProof != nil {
// 		kbAddress := domain.NewPartKYCBuilder()
// 		kbAddress.WithBasics(record.AddressProof.DocumentID, record.AddressProof.DocType, record.AddressProof.ModifiedBy)
// 		kbAddress.Status(record.AddressProof.Status)
// 		profileAndKYC.AddressProof = kbAddress.Build()
// 	}
// 	if record.IDProof != nil {
// 		kbID := domain.NewPartKYCBuilder()
// 		kbID.WithBasics(record.IDProof.DocumentID, record.IDProof.DocType, record.IDProof.ModifiedBy)
// 		kbID.Status(record.IDProof.Status)
// 		profileAndKYC.IdProof = kbID.Build()
// 	}
// 	if record.PanProof != nil {
// 		kbPan := domain.NewPartKYCBuilder()
// 		kbPan.WithBasics(record.PanProof.DocumentID, record.PanProof.DocType, record.PanProof.ModifiedBy)
// 		kbPan.Status(record.PanProof.Status)
// 		profileAndKYC.PANProof = kbPan.Build()
// 	}
// 	return profileAndKYC
// }

// // TODO need to remove this just a temp solution
// func (k KYC) GetJWRToekn() string {
// 	return k.jwrToken
// }
