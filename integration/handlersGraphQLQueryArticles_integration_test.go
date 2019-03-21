// +build integration

package integration

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"testing"
	"time"

	"github.com/go-test/deep"
)

func TestHandlersGraphQLQueryArticles(t *testing.T) {
	ts := newTserver()
	defer ts.closeAll()
	testCases := []struct {
		description string
		body        map[string]interface{}
		query       string
		expectBody  map[string]interface{}
	}{
		{
			description: "testing normal SG + EN_US case",
			body: map[string]interface{}{
				"query": `
				{
					allArticles {
						page
						perPage
						pageCount
						count
						articles {
							id
							authorId
							commentsDisable
							draft
							promoted
							position
							voteSum
							voteCount
							createdAt
							updatedAt
							sourceLocale
							outdated
							outdatedLocales
							editedAt
							labelNames
							countryCode
							url
							htmlUrl
							name
							title
							body
							locale
						}
					}
				}
				`,
			},
			expectBody: map[string]interface{}{
				"data": map[string]interface{}{
					"allArticles": map[string]interface{}{
						"page":      1,
						"perPage":   30,
						"pageCount": 1,
						"count":     4,
						"articles": []interface{}{
							map[string]interface{}{
								"id":              "115016053147",
								"authorId":        "7222048487",
								"commentsDisable": false,
								"draft":           false,
								"promoted":        false,
								"position":        0,
								"voteSum":         17,
								"voteCount":       33,
								"createdAt":       time.Date(2018, 1, 12, 3, 24, 59, 0, time.UTC),
								"updatedAt":       time.Date(2018, 7, 16, 13, 35, 30, 0, time.UTC),
								"sourceLocale":    "en-us",
								"outdated":        false,
								"outdatedLocales": []string{},
								"editedAt":        time.Date(2018, 7, 10, 4, 17, 18, 0, time.UTC),
								"labelNames":      []string{"use_for_answer_bot", "preparing", "ontheway", "confirmed"},
								"countryCode":     "sg",
								"url":             "https://honestbeehelp-sg.zendesk.com/api/v2/help_center/en-us/articles/115016053147-How-do-I-prepare-my-laundry-for-pickup-.json",
								"htmlUrl":         "https://honestbeehelp-sg.zendesk.com/hc/en-us/articles/115016053147-How-do-I-prepare-my-laundry-for-pickup-",
								"name":            "How do I prepare my laundry for pickup?",
								"title":           "How do I prepare my laundry for pickup?",
								"body":            "<p>Please be ready with your laundry when the delivery bee comes. Also check that all clothing pockets are empty, as it will be difficult for us to trace personal articles after your laundry has been handed over. We will not be liable for missing personal articles after you hand it over to us as well.</p>",
								"locale":          "en-us",
							},
							map[string]interface{}{
								"id":              "115016039687",
								"authorId":        "7222048487",
								"commentsDisable": false,
								"draft":           false,
								"promoted":        false,
								"position":        0,
								"voteSum":         -1,
								"voteCount":       9,
								"createdAt":       time.Date(2018, 1, 11, 11, 49, 48, 0, time.UTC),
								"updatedAt":       time.Date(2018, 7, 10, 4, 17, 49, 0, time.UTC),
								"sourceLocale":    "en-us",
								"outdated":        false,
								"outdatedLocales": []string{},
								"editedAt":        time.Date(2018, 1, 11, 11, 49, 48, 0, time.UTC),
								"labelNames":      []string{"delivered"},
								"countryCode":     "sg",
								"url":             "https://honestbeehelp-sg.zendesk.com/api/v2/help_center/en-us/articles/115016039687-I-have-not-received-my-order-confirmation-final-receipt.json",
								"htmlUrl":         "https://honestbeehelp-sg.zendesk.com/hc/en-us/articles/115016039687-I-have-not-received-my-order-confirmation-final-receipt",
								"name":            "I have not received my order confirmation/final receipt",
								"title":           "I have not received my order confirmation/final receipt",
								"body":            "<p>We send you your order confirmation immediately after you have placed an order. There is a chance that it might have ended up in the junk/spam folder. Please let us know if you do not find it there.</p>\\n<p><br>We aim to send you your final receipt within 5 working days after we have returned your laundry. Please note that we might adjust your final receipt for any additional articles added to the order and the difference in weight (if any) of your wash and fold loads.</p>\\n<p><br>If this answer did not solve your issue, please contact us and we will do our best to make things right for you. </p>",
								"locale":          "en-us",
							},
							map[string]interface{}{
								"id":              "115015447167",
								"authorId":        "24400224208",
								"commentsDisable": false,
								"draft":           false,
								"promoted":        false,
								"position":        0,
								"voteSum":         0,
								"voteCount":       -30,
								"createdAt":       time.Date(2017, 11, 29, 2, 37, 50, 0, time.UTC),
								"updatedAt":       time.Date(2018, 7, 14, 10, 33, 44, 0, time.UTC),
								"sourceLocale":    "en-us",
								"outdated":        false,
								"outdatedLocales": []string{},
								"editedAt":        time.Date(2018, 7, 9, 9, 49, 23, 0, time.UTC),
								"labelNames":      []string{"preparing", "confirmed"},
								"countryCode":     "sg",
								"url":             "https://honestbeehelp-sg.zendesk.com/api/v2/help_center/en-us/articles/115015447167-I-want-to-cancel-my-food-order.json",
								"htmlUrl":         "https://honestbeehelp-sg.zendesk.com/hc/en-us/articles/115015447167-I-want-to-cancel-my-food-order",
								"name":            "I want to cancel my food order",
								"title":           "I want to cancel my food order",
								"body":            "<p><strong>Can I cancel my food order? </strong></p>\\n<p>You may cancel your order before the restaurant has started preparing it. Simply go to <a href=\"https://www.honestbee.sg/en/food/orders\">Your Orders</a> and select the order you wish to cancel.</p>\\n<p>If you do not want your order to be delivered but the food is being prepared or our deliverer is on his way to you, please contact us for assistance. Do note that charges of the food order may still apply to you.</p>\\n<p> </p>\\n<p><button><a href=\"javascript:zE.show();zE.activate();\">Chat with Us</a></button></p>",
								"locale":          "en-us",
							},
							map[string]interface{}{
								"id":              "115015433907",
								"authorId":        "24400224208",
								"commentsDisable": false,
								"draft":           false,
								"promoted":        false,
								"position":        2,
								"voteSum":         -12,
								"voteCount":       16,
								"createdAt":       time.Date(2017, 11, 28, 13, 2, 47, 0, time.UTC),
								"updatedAt":       time.Date(2018, 7, 17, 11, 3, 40, 0, time.UTC),
								"sourceLocale":    "en-us",
								"outdated":        false,
								"outdatedLocales": []string{},
								"editedAt":        time.Date(2018, 7, 10, 4, 26, 17, 0, time.UTC),
								"labelNames":      []string{"ontheway", "delivered"},
								"countryCode":     "sg",
								"url":             "https://honestbeehelp-sg.zendesk.com/api/v2/help_center/en-us/articles/115015433907-What-if-I-am-not-at-home-to-receive-my-order-.json",
								"htmlUrl":         "https://honestbeehelp-sg.zendesk.com/hc/en-us/articles/115015433907-What-if-I-am-not-at-home-to-receive-my-order-",
								"name":            "What if I am not at home to receive my order?",
								"title":           "What if I am not at home to receive my order?",
								"body":            "<p>If you missed your delivery, please contact us as soon as possible. We will do our best to reschedule your delivery. We reserve the right to charge an additional redelivery fee of $15. Perishables such as meat and frozen food may be discarded in accordance with Food Safety Standards, and are not refundable on your order. </p>\\n<p> </p>\\n<p><button><a href=\"javascript:zE.show();zE.activate();\">Chat with Us</a></button></p>",
								"locale":          "en-us",
							},
						},
					},
				},
			},
		},
		{
			description: "testing normal TW + EN_US case",
			body: map[string]interface{}{
				"query": `
				{
					allArticles(countryCode: TW, locale: EN_US) {
						page
						perPage
						pageCount
						count
						articles
						{
							id
							authorId
							commentsDisable
							draft
							promoted
							position
							voteSum
							voteCount
							createdAt
							updatedAt
							sourceLocale
							outdated
							outdatedLocales
							editedAt
							labelNames
							countryCode
							url
							htmlUrl
							name
							title
							body
							locale
							categoryConnection {
								id
								name
								keyName
							}
							sectionConnection {
								id
								name
							}
						}
					}
				}
				`,
			},
			expectBody: map[string]interface{}{
				"data": map[string]interface{}{
					"allArticles": map[string]interface{}{
						"page":      1,
						"perPage":   30,
						"pageCount": 1,
						"count":     5,
						"articles": []interface{}{
							map[string]interface{}{
								"id":              "115015959188",
								"authorId":        "24400224208",
								"commentsDisable": false,
								"draft":           false,
								"promoted":        false,
								"position":        0,
								"voteSum":         0,
								"voteCount":       0,
								"createdAt":       time.Date(2017, 12, 27, 3, 14, 16, 0, time.UTC),
								"updatedAt":       time.Date(2017, 12, 27, 3, 14, 48, 0, time.UTC),
								"sourceLocale":    "zh-tw",
								"outdated":        false,
								"outdatedLocales": []string{},
								"editedAt":        time.Date(2017, 12, 27, 3, 14, 16, 0, time.UTC),
								"labelNames":      []string{},
								"countryCode":     "tw",
								"url":             "https://honestbeehelp-tw.zendesk.com/api/v2/help_center/en-us/articles/115015959188-What-can-I-do-when-my-cart-is-locked-.json",
								"htmlUrl":         "https://help.honestbee.tw/hc/en-us/articles/115015959188-What-can-I-do-when-my-cart-is-locked-",
								"name":            "What can I do when my cart is locked?",
								"title":           "What can I do when my cart is locked?",
								"body":            "<p>When there is an error completing your checkout,your cart will be temporarily locked to prevent further changes to your order. To unlock your cart,click ‘Yes,unlock my cart,’ when prompted.</p>",
								"locale":          "en-us",
								"categoryConnection": map[string]interface{}{
									"id":      "115002432448",
									"name":    "My Account",
									"keyName": "myAccount",
								},
								"sectionConnection": map[string]interface{}{
									"id":   "115004118448",
									"name": "I need help with my account",
								},
							},
							map[string]interface{}{
								"id":              "115015885547",
								"authorId":        "24400224208",
								"commentsDisable": false,
								"draft":           false,
								"promoted":        false,
								"position":        0,
								"voteSum":         0,
								"voteCount":       0,
								"createdAt":       time.Date(2017, 12, 27, 3, 12, 58, 0, time.UTC),
								"updatedAt":       time.Date(2017, 12, 27, 3, 13, 42, 0, time.UTC),
								"sourceLocale":    "zh-tw",
								"outdated":        false,
								"outdatedLocales": []string{},
								"editedAt":        time.Date(2017, 12, 27, 3, 12, 58, 0, time.UTC),
								"labelNames":      []string{},
								"countryCode":     "tw",
								"url":             "https://honestbeehelp-tw.zendesk.com/api/v2/help_center/en-us/articles/115015885547-What-is-honestbee-s-information-security-policy-.json",
								"htmlUrl":         "https://help.honestbee.tw/hc/en-us/articles/115015885547-What-is-honestbee-s-information-security-policy-",
								"name":            "What is honestbee’s information security policy?",
								"title":           "What is honestbee’s information security policy?",
								"body":            `<p>honestbee takes your privacy seriously and complies with all the relevant laws to ensure your details are kept secure. Read our <a href="https://www.honestbee.tw/privacy-policy">Privacy Policy</a> for more information.</p>`,
								"locale":          "en-us",
								"categoryConnection": map[string]interface{}{
									"id":      "115002432448",
									"name":    "My Account",
									"keyName": "myAccount",
								},
								"sectionConnection": map[string]interface{}{
									"id":   "115004118448",
									"name": "I need help with my account",
								},
							},
							map[string]interface{}{
								"id":              "115015885507",
								"authorId":        "24400224208",
								"commentsDisable": false,
								"draft":           false,
								"promoted":        false,
								"position":        0,
								"voteSum":         0,
								"voteCount":       0,
								"createdAt":       time.Date(2017, 12, 27, 3, 11, 30, 0, time.UTC),
								"updatedAt":       time.Date(2018, 1, 10, 6, 7, 41, 0, time.UTC),
								"sourceLocale":    "zh-tw",
								"outdated":        false,
								"outdatedLocales": []string{},
								"editedAt":        time.Date(2018, 1, 10, 6, 7, 41, 0, time.UTC),
								"labelNames":      []string{},
								"countryCode":     "tw",
								"url":             "https://honestbeehelp-tw.zendesk.com/api/v2/help_center/en-us/articles/115015885507-How-can-I-change-my-email-address-.json",
								"htmlUrl":         "https://help.honestbee.tw/hc/en-us/articles/115015885507-How-can-I-change-my-email-address-",
								"name":            "How can I change my email address?",
								"title":           "How can I change my email address?",
								"body":            `<p>Currently,it’s not possible to change your email address. To register with a different email address,please create a new account.</p>`,
								"locale":          "en-us",
								"categoryConnection": map[string]interface{}{
									"id":      "115002432448",
									"name":    "My Account",
									"keyName": "myAccount",
								},
								"sectionConnection": map[string]interface{}{
									"id":   "115004118448",
									"name": "I need help with my account",
								},
							},
							map[string]interface{}{
								"id":              "115015959168",
								"authorId":        "24400224208",
								"commentsDisable": false,
								"draft":           false,
								"promoted":        false,
								"position":        0,
								"voteSum":         -2,
								"voteCount":       -2,
								"createdAt":       time.Date(2017, 12, 27, 3, 8, 5, 0, time.UTC),
								"updatedAt":       time.Date(2018, 3, 6, 12, 39, 30, 0, time.UTC),
								"sourceLocale":    "zh-tw",
								"outdated":        false,
								"outdatedLocales": []string{},
								"editedAt":        time.Date(2018, 1, 10, 6, 7, 41, 0, time.UTC),
								"labelNames":      []string{},
								"countryCode":     "tw",
								"url":             "https://honestbeehelp-tw.zendesk.com/api/v2/help_center/en-us/articles/115015959168-How-can-I-change-my-credit-card-password-or-delivery-address-.json",
								"htmlUrl":         "https://help.honestbee.tw/hc/en-us/articles/115015959168-How-can-I-change-my-credit-card-password-or-delivery-address-",
								"name":            "How can I change my credit card,password or delivery address?",
								"title":           "How can I change my credit card,password or delivery address?",
								"body":            `<p>Log in to your account and go to your profile icon at the top right corner. Select Settings from the dropdown menu to edit your details.</p>`,
								"locale":          "en-us",
								"categoryConnection": map[string]interface{}{
									"id":      "115002432448",
									"name":    "My Account",
									"keyName": "myAccount",
								},
								"sectionConnection": map[string]interface{}{
									"id":   "115004118448",
									"name": "I need help with my account",
								},
							},
							map[string]interface{}{
								"id":              "115015959148",
								"authorId":        "24400224208",
								"commentsDisable": false,
								"draft":           false,
								"promoted":        false,
								"position":        0,
								"voteSum":         -1,
								"voteCount":       1,
								"createdAt":       time.Date(2017, 12, 27, 2, 59, 48, 0, time.UTC),
								"updatedAt":       time.Date(2018, 3, 2, 5, 59, 58, 0, time.UTC),
								"sourceLocale":    "zh-tw",
								"outdated":        false,
								"outdatedLocales": []string{},
								"editedAt":        time.Date(2017, 12, 27, 3, 6, 54, 0, time.UTC),
								"labelNames":      []string{},
								"countryCode":     "tw",
								"url":             "https://honestbeehelp-tw.zendesk.com/api/v2/help_center/en-us/articles/115015959148-Help-I-ve-forgotten-my-password-.json",
								"htmlUrl":         "https://help.honestbee.tw/hc/en-us/articles/115015959148-Help-I-ve-forgotten-my-password-",
								"name":            "Help! I’ve forgotten my password.",
								"title":           "Help! I’ve forgotten my password.",
								"body":            "<p>Click on the Forgot Password link on the Login page and enter your registered email address. We’ll send you an email to reset your password. Occasionally emails end up in the junk/spam folder. Take a look there.</p>",
								"locale":          "en-us",
								"categoryConnection": map[string]interface{}{
									"id":      "115002432448",
									"name":    "My Account",
									"keyName": "myAccount",
								},
								"sectionConnection": map[string]interface{}{
									"id":   "115004118448",
									"name": "I need help with my account",
								},
							},
						},
					},
				},
			},
		},
		{
			description: "testing normal TW + ZH_TW case",
			body: map[string]interface{}{
				"query": `
				{
					allArticles(countryCode: TW, locale: ZH_TW) {
						page
						perPage
						pageCount
						count
						articles {
							id
							authorId
							commentsDisable
							draft
							promoted
							position
							voteSum
							voteCount
							createdAt
							updatedAt
							sourceLocale
							outdated
							outdatedLocales
							editedAt
							labelNames
							countryCode
							url
							htmlUrl
							name
							title
							body
							locale
							categoryConnection {
								id
								name
								keyName
							}
							sectionConnection {
								id
								name
							}
						}
					}
				}
				`,
			},
			expectBody: map[string]interface{}{
				"data": map[string]interface{}{
					"allArticles": map[string]interface{}{
						"page":      1,
						"perPage":   30,
						"pageCount": 1,
						"count":     5,
						"articles": []interface{}{
							map[string]interface{}{
								"id":              "115015959188",
								"authorId":        "24400224208",
								"commentsDisable": false,
								"draft":           false,
								"promoted":        false,
								"position":        0,
								"voteSum":         0,
								"voteCount":       0,
								"createdAt":       time.Date(2017, 12, 27, 3, 14, 16, 0, time.UTC),
								"updatedAt":       time.Date(2017, 12, 27, 3, 14, 48, 0, time.UTC),
								"sourceLocale":    "zh-tw",
								"outdated":        false,
								"outdatedLocales": []string{},
								"editedAt":        time.Date(2017, 12, 27, 3, 14, 16, 0, time.UTC),
								"labelNames":      []string{},
								"countryCode":     "tw",
								"url":             "https://honestbeehelp-tw.zendesk.com/api/v2/help_center/zh-tw/articles/115015959188-%E6%88%91%E7%9A%84%E8%B3%BC%E7%89%A9%E8%BB%8A%E9%8E%96%E4%BD%8F%E4%BA%86-%E8%A9%B2%E6%80%8E%E9%BA%BC%E8%BE%A6-.json",
								"htmlUrl":         "https://help.honestbee.tw/hc/zh-tw/articles/115015959188-%E6%88%91%E7%9A%84%E8%B3%BC%E7%89%A9%E8%BB%8A%E9%8E%96%E4%BD%8F%E4%BA%86-%E8%A9%B2%E6%80%8E%E9%BA%BC%E8%BE%A6-",
								"name":            "我的購物車鎖住了，該怎麼辦？",
								"title":           "我的購物車鎖住了，該怎麼辦？",
								"body":            "<p>當結帳出現錯誤時，您的購物車會暫時被鎖住，以避免您的訂單出現異動。要解鎖您的購物車，請在出現提示時，點選「是的，解鎖我的購物車」。</p>",
								"locale":          "zh-tw",
								"categoryConnection": map[string]interface{}{
									"id":      "115002432448",
									"name":    "My Account",
									"keyName": "myAccount",
								},
								"sectionConnection": map[string]interface{}{
									"id":   "115004118448",
									"name": "I need help with my account",
								},
							},
							map[string]interface{}{
								"id":              "115015885547",
								"authorId":        "24400224208",
								"commentsDisable": false,
								"draft":           false,
								"promoted":        false,
								"position":        0,
								"voteSum":         0,
								"voteCount":       0,
								"createdAt":       time.Date(2017, 12, 27, 3, 12, 58, 0, time.UTC),
								"updatedAt":       time.Date(2017, 12, 27, 3, 13, 42, 0, time.UTC),
								"sourceLocale":    "zh-tw",
								"outdated":        false,
								"outdatedLocales": []string{},
								"editedAt":        time.Date(2017, 12, 27, 3, 12, 58, 0, time.UTC),
								"labelNames":      []string{},
								"countryCode":     "tw",
								"url":             "https://honestbeehelp-tw.zendesk.com/api/v2/help_center/zh-tw/articles/115015885547-honestbee-%E7%9A%84%E5%AE%89%E5%85%A8%E6%94%BF%E7%AD%96%E7%82%BA%E4%BD%95-.json",
								"htmlUrl":         "https://help.honestbee.tw/hc/zh-tw/articles/115015885547-honestbee-%E7%9A%84%E5%AE%89%E5%85%A8%E6%94%BF%E7%AD%96%E7%82%BA%E4%BD%95-",
								"name":            "honestbee 的安全政策為何？",
								"title":           "honestbee 的安全政策為何？",
								"body":            `<p>honestbee 很重視您的隱私，並遵循所有相關法規以確保您的資訊安全。請閱讀我們的<a href="https://www.honestbee.tw/privacy-policy">隱私權政策</a>以了解更多資訊。</p>`,
								"locale":          "zh-tw",
								"categoryConnection": map[string]interface{}{
									"id":      "115002432448",
									"name":    "My Account",
									"keyName": "myAccount",
								},
								"sectionConnection": map[string]interface{}{
									"id":   "115004118448",
									"name": "I need help with my account",
								},
							},
							map[string]interface{}{
								"id":              "115015885507",
								"authorId":        "24400224208",
								"commentsDisable": false,
								"draft":           false,
								"promoted":        false,
								"position":        0,
								"voteSum":         0,
								"voteCount":       0,
								"createdAt":       time.Date(2017, 12, 27, 3, 11, 30, 0, time.UTC),
								"updatedAt":       time.Date(2018, 1, 10, 6, 7, 41, 0, time.UTC),
								"sourceLocale":    "zh-tw",
								"outdated":        false,
								"outdatedLocales": []string{},
								"editedAt":        time.Date(2018, 1, 10, 6, 7, 41, 0, time.UTC),
								"labelNames":      []string{},
								"countryCode":     "tw",
								"url":             "https://honestbeehelp-tw.zendesk.com/api/v2/help_center/zh-tw/articles/115015885507-%E5%A6%82%E4%BD%95%E6%9B%B4%E6%94%B9%E9%9B%BB%E5%AD%90%E4%BF%A1%E7%AE%B1-.json",
								"htmlUrl":         "https://help.honestbee.tw/hc/zh-tw/articles/115015885507-%E5%A6%82%E4%BD%95%E6%9B%B4%E6%94%B9%E9%9B%BB%E5%AD%90%E4%BF%A1%E7%AE%B1-",
								"name":            "如何更改電子信箱？",
								"title":           "如何更改電子信箱？",
								"body":            `<p>您的電子信箱目前無法更改。若要以不同信箱進行註冊，請建立新帳號。</p>`,
								"locale":          "zh-tw",
								"categoryConnection": map[string]interface{}{
									"id":      "115002432448",
									"name":    "My Account",
									"keyName": "myAccount",
								},
								"sectionConnection": map[string]interface{}{
									"id":   "115004118448",
									"name": "I need help with my account",
								},
							},
							map[string]interface{}{
								"id":              "115015959168",
								"authorId":        "24400224208",
								"commentsDisable": false,
								"draft":           false,
								"promoted":        false,
								"position":        0,
								"voteSum":         -2,
								"voteCount":       -2,
								"createdAt":       time.Date(2017, 12, 27, 3, 8, 5, 0, time.UTC),
								"updatedAt":       time.Date(2018, 3, 6, 12, 39, 30, 0, time.UTC),
								"sourceLocale":    "zh-tw",
								"outdated":        false,
								"outdatedLocales": []string{},
								"editedAt":        time.Date(2018, 1, 10, 6, 7, 41, 0, time.UTC),
								"labelNames":      []string{},
								"countryCode":     "tw",
								"url":             "https://honestbeehelp-tw.zendesk.com/api/v2/help_center/zh-tw/articles/115015959168-%E5%A6%82%E4%BD%95%E6%9B%B4%E6%94%B9%E4%BF%A1%E7%94%A8%E5%8D%A1-%E5%AF%86%E7%A2%BC%E6%88%96%E9%80%81%E8%B2%A8%E5%9C%B0%E5%9D%80-.json",
								"htmlUrl":         "https://help.honestbee.tw/hc/zh-tw/articles/115015959168-%E5%A6%82%E4%BD%95%E6%9B%B4%E6%94%B9%E4%BF%A1%E7%94%A8%E5%8D%A1-%E5%AF%86%E7%A2%BC%E6%88%96%E9%80%81%E8%B2%A8%E5%9C%B0%E5%9D%80-",
								"name":            "如何更改信用卡、密碼或送貨地址？",
								"title":           "如何更改信用卡、密碼或送貨地址？",
								"body":            `<p>登入您的帳號並前往右上角的個人資料圖示。在下拉式選單中選擇「設定」來編輯您的資料。</p>`,
								"locale":          "zh-tw",
								"categoryConnection": map[string]interface{}{
									"id":      "115002432448",
									"name":    "My Account",
									"keyName": "myAccount",
								},
								"sectionConnection": map[string]interface{}{
									"id":   "115004118448",
									"name": "I need help with my account",
								},
							},
							map[string]interface{}{
								"id":              "115015959148",
								"authorId":        "24400224208",
								"commentsDisable": false,
								"draft":           false,
								"promoted":        false,
								"position":        0,
								"voteSum":         -1,
								"voteCount":       1,
								"createdAt":       time.Date(2017, 12, 27, 2, 59, 48, 0, time.UTC),
								"updatedAt":       time.Date(2018, 3, 2, 5, 59, 58, 0, time.UTC),
								"sourceLocale":    "zh-tw",
								"outdated":        false,
								"outdatedLocales": []string{},
								"editedAt":        time.Date(2017, 12, 27, 3, 6, 54, 0, time.UTC),
								"labelNames":      []string{},
								"countryCode":     "tw",
								"url":             "https://honestbeehelp-tw.zendesk.com/api/v2/help_center/zh-tw/articles/115015959148-%E6%95%91%E5%91%BD-%E6%88%91%E5%BF%98%E8%A8%98%E5%AF%86%E7%A2%BC%E4%BA%86.json",
								"htmlUrl":         "https://help.honestbee.tw/hc/zh-tw/articles/115015959148-%E6%95%91%E5%91%BD-%E6%88%91%E5%BF%98%E8%A8%98%E5%AF%86%E7%A2%BC%E4%BA%86",
								"name":            "救命！我忘記密碼了",
								"title":           "救命！我忘記密碼了",
								"body":            "<p>請在登入頁面點選「忘記密碼」，並輸入您註冊時所使用的電子信箱。我們會寄給您一封信讓您重設密碼。有時候信件會跑到垃圾信件匣，請務必檢查看看。</p>",
								"locale":          "zh-tw",
								"categoryConnection": map[string]interface{}{
									"id":      "115002432448",
									"name":    "My Account",
									"keyName": "myAccount",
								},
								"sectionConnection": map[string]interface{}{
									"id":   "115004118448",
									"name": "I need help with my account",
								},
							},
						},
					},
				},
			},
		},
		{
			description: "testing sort by created at case",
			body: map[string]interface{}{
				"query": `
				{
					allArticles(countryCode: TW, locale: EN_US, sortBy: CREATED_AT) {
						page
						perPage
						pageCount
						count
						articles {
							id
							authorId
							commentsDisable
							draft
							promoted
							position
							voteSum
							voteCount
							createdAt
							updatedAt
							sourceLocale
							outdated
							outdatedLocales
							editedAt
							labelNames
							countryCode
							url
							htmlUrl
							name
							title
							body
							locale
							categoryConnection {
								id
								name
								keyName
							}
							sectionConnection {
								id
								name
							}
						}
					}
				}
				`,
			},
			expectBody: map[string]interface{}{
				"data": map[string]interface{}{
					"allArticles": map[string]interface{}{
						"page":      1,
						"perPage":   30,
						"pageCount": 1,
						"count":     5,
						"articles": []interface{}{
							map[string]interface{}{
								"id":              "115015959148",
								"authorId":        "24400224208",
								"commentsDisable": false,
								"draft":           false,
								"promoted":        false,
								"position":        0,
								"voteSum":         -1,
								"voteCount":       1,
								"createdAt":       time.Date(2017, 12, 27, 2, 59, 48, 0, time.UTC),
								"updatedAt":       time.Date(2018, 3, 2, 5, 59, 58, 0, time.UTC),
								"sourceLocale":    "zh-tw",
								"outdated":        false,
								"outdatedLocales": []string{},
								"editedAt":        time.Date(2017, 12, 27, 3, 6, 54, 0, time.UTC),
								"labelNames":      []string{},
								"countryCode":     "tw",
								"url":             "https://honestbeehelp-tw.zendesk.com/api/v2/help_center/en-us/articles/115015959148-Help-I-ve-forgotten-my-password-.json",
								"htmlUrl":         "https://help.honestbee.tw/hc/en-us/articles/115015959148-Help-I-ve-forgotten-my-password-",
								"name":            "Help! I’ve forgotten my password.",
								"title":           "Help! I’ve forgotten my password.",
								"body":            "<p>Click on the Forgot Password link on the Login page and enter your registered email address. We’ll send you an email to reset your password. Occasionally emails end up in the junk/spam folder. Take a look there.</p>",
								"locale":          "en-us",
								"categoryConnection": map[string]interface{}{
									"id":      "115002432448",
									"name":    "My Account",
									"keyName": "myAccount",
								},
								"sectionConnection": map[string]interface{}{
									"id":   "115004118448",
									"name": "I need help with my account",
								},
							},
							map[string]interface{}{
								"id":              "115015959168",
								"authorId":        "24400224208",
								"commentsDisable": false,
								"draft":           false,
								"promoted":        false,
								"position":        0,
								"voteSum":         -2,
								"voteCount":       -2,
								"createdAt":       time.Date(2017, 12, 27, 3, 8, 5, 0, time.UTC),
								"updatedAt":       time.Date(2018, 3, 6, 12, 39, 30, 0, time.UTC),
								"sourceLocale":    "zh-tw",
								"outdated":        false,
								"outdatedLocales": []string{},
								"editedAt":        time.Date(2018, 1, 10, 6, 7, 41, 0, time.UTC),
								"labelNames":      []string{},
								"countryCode":     "tw",
								"url":             "https://honestbeehelp-tw.zendesk.com/api/v2/help_center/en-us/articles/115015959168-How-can-I-change-my-credit-card-password-or-delivery-address-.json",
								"htmlUrl":         "https://help.honestbee.tw/hc/en-us/articles/115015959168-How-can-I-change-my-credit-card-password-or-delivery-address-",
								"name":            "How can I change my credit card,password or delivery address?",
								"title":           "How can I change my credit card,password or delivery address?",
								"body":            `<p>Log in to your account and go to your profile icon at the top right corner. Select Settings from the dropdown menu to edit your details.</p>`,
								"locale":          "en-us",
								"categoryConnection": map[string]interface{}{
									"id":      "115002432448",
									"name":    "My Account",
									"keyName": "myAccount",
								},
								"sectionConnection": map[string]interface{}{
									"id":   "115004118448",
									"name": "I need help with my account",
								},
							},
							map[string]interface{}{
								"id":              "115015885507",
								"authorId":        "24400224208",
								"commentsDisable": false,
								"draft":           false,
								"promoted":        false,
								"position":        0,
								"voteSum":         0,
								"voteCount":       0,
								"createdAt":       time.Date(2017, 12, 27, 3, 11, 30, 0, time.UTC),
								"updatedAt":       time.Date(2018, 1, 10, 6, 7, 41, 0, time.UTC),
								"sourceLocale":    "zh-tw",
								"outdated":        false,
								"outdatedLocales": []string{},
								"editedAt":        time.Date(2018, 1, 10, 6, 7, 41, 0, time.UTC),
								"labelNames":      []string{},
								"countryCode":     "tw",
								"url":             "https://honestbeehelp-tw.zendesk.com/api/v2/help_center/en-us/articles/115015885507-How-can-I-change-my-email-address-.json",
								"htmlUrl":         "https://help.honestbee.tw/hc/en-us/articles/115015885507-How-can-I-change-my-email-address-",
								"name":            "How can I change my email address?",
								"title":           "How can I change my email address?",
								"body":            `<p>Currently,it’s not possible to change your email address. To register with a different email address,please create a new account.</p>`,
								"locale":          "en-us",
								"categoryConnection": map[string]interface{}{
									"id":      "115002432448",
									"name":    "My Account",
									"keyName": "myAccount",
								},
								"sectionConnection": map[string]interface{}{
									"id":   "115004118448",
									"name": "I need help with my account",
								},
							},
							map[string]interface{}{
								"id":              "115015885547",
								"authorId":        "24400224208",
								"commentsDisable": false,
								"draft":           false,
								"promoted":        false,
								"position":        0,
								"voteSum":         0,
								"voteCount":       0,
								"createdAt":       time.Date(2017, 12, 27, 3, 12, 58, 0, time.UTC),
								"updatedAt":       time.Date(2017, 12, 27, 3, 13, 42, 0, time.UTC),
								"sourceLocale":    "zh-tw",
								"outdated":        false,
								"outdatedLocales": []string{},
								"editedAt":        time.Date(2017, 12, 27, 3, 12, 58, 0, time.UTC),
								"labelNames":      []string{},
								"countryCode":     "tw",
								"url":             "https://honestbeehelp-tw.zendesk.com/api/v2/help_center/en-us/articles/115015885547-What-is-honestbee-s-information-security-policy-.json",
								"htmlUrl":         "https://help.honestbee.tw/hc/en-us/articles/115015885547-What-is-honestbee-s-information-security-policy-",
								"name":            "What is honestbee’s information security policy?",
								"title":           "What is honestbee’s information security policy?",
								"body":            `<p>honestbee takes your privacy seriously and complies with all the relevant laws to ensure your details are kept secure. Read our <a href="https://www.honestbee.tw/privacy-policy">Privacy Policy</a> for more information.</p>`,
								"locale":          "en-us",
								"categoryConnection": map[string]interface{}{
									"id":      "115002432448",
									"name":    "My Account",
									"keyName": "myAccount",
								},
								"sectionConnection": map[string]interface{}{
									"id":   "115004118448",
									"name": "I need help with my account",
								},
							},
							map[string]interface{}{
								"id":              "115015959188",
								"authorId":        "24400224208",
								"commentsDisable": false,
								"draft":           false,
								"promoted":        false,
								"position":        0,
								"voteSum":         0,
								"voteCount":       0,
								"createdAt":       time.Date(2017, 12, 27, 3, 14, 16, 0, time.UTC),
								"updatedAt":       time.Date(2017, 12, 27, 3, 14, 48, 0, time.UTC),
								"sourceLocale":    "zh-tw",
								"outdated":        false,
								"outdatedLocales": []string{},
								"editedAt":        time.Date(2017, 12, 27, 3, 14, 16, 0, time.UTC),
								"labelNames":      []string{},
								"countryCode":     "tw",
								"url":             "https://honestbeehelp-tw.zendesk.com/api/v2/help_center/en-us/articles/115015959188-What-can-I-do-when-my-cart-is-locked-.json",
								"htmlUrl":         "https://help.honestbee.tw/hc/en-us/articles/115015959188-What-can-I-do-when-my-cart-is-locked-",
								"name":            "What can I do when my cart is locked?",
								"title":           "What can I do when my cart is locked?",
								"body":            "<p>When there is an error completing your checkout,your cart will be temporarily locked to prevent further changes to your order. To unlock your cart,click ‘Yes,unlock my cart,’ when prompted.</p>",
								"locale":          "en-us",
								"categoryConnection": map[string]interface{}{
									"id":      "115002432448",
									"name":    "My Account",
									"keyName": "myAccount",
								},
								"sectionConnection": map[string]interface{}{
									"id":   "115004118448",
									"name": "I need help with my account",
								},
							},
						},
					},
				},
			},
		},
		{
			description: "testing sort by updated at case",
			body: map[string]interface{}{
				"query": `
				{
					allArticles(countryCode: TW, locale: EN_US, sortBy: UPDATED_AT) {
						page
						perPage
						pageCount
						count
						articles {
							id
							authorId
							commentsDisable
							draft
							promoted
							position
							voteSum
							voteCount
							createdAt
							updatedAt
							sourceLocale
							outdated
							outdatedLocales
							editedAt
							labelNames
							countryCode
							url
							htmlUrl
							name
							title
							body
							locale
							categoryConnection {
								id
								name
								keyName
							}
							sectionConnection {
								id
								name
							}
						}
					}
				}
				`,
			},
			expectBody: map[string]interface{}{
				"data": map[string]interface{}{
					"allArticles": map[string]interface{}{
						"page":      1,
						"perPage":   30,
						"pageCount": 1,
						"count":     5,
						"articles": []interface{}{
							map[string]interface{}{
								"id":              "115015885547",
								"authorId":        "24400224208",
								"commentsDisable": false,
								"draft":           false,
								"promoted":        false,
								"position":        0,
								"voteSum":         0,
								"voteCount":       0,
								"createdAt":       time.Date(2017, 12, 27, 3, 12, 58, 0, time.UTC),
								"updatedAt":       time.Date(2017, 12, 27, 3, 13, 42, 0, time.UTC),
								"sourceLocale":    "zh-tw",
								"outdated":        false,
								"outdatedLocales": []string{},
								"editedAt":        time.Date(2017, 12, 27, 3, 12, 58, 0, time.UTC),
								"labelNames":      []string{},
								"countryCode":     "tw",
								"url":             "https://honestbeehelp-tw.zendesk.com/api/v2/help_center/en-us/articles/115015885547-What-is-honestbee-s-information-security-policy-.json",
								"htmlUrl":         "https://help.honestbee.tw/hc/en-us/articles/115015885547-What-is-honestbee-s-information-security-policy-",
								"name":            "What is honestbee’s information security policy?",
								"title":           "What is honestbee’s information security policy?",
								"body":            `<p>honestbee takes your privacy seriously and complies with all the relevant laws to ensure your details are kept secure. Read our <a href="https://www.honestbee.tw/privacy-policy">Privacy Policy</a> for more information.</p>`,
								"locale":          "en-us",
								"categoryConnection": map[string]interface{}{
									"id":      "115002432448",
									"name":    "My Account",
									"keyName": "myAccount",
								},
								"sectionConnection": map[string]interface{}{
									"id":   "115004118448",
									"name": "I need help with my account",
								},
							},
							map[string]interface{}{
								"id":              "115015959188",
								"authorId":        "24400224208",
								"commentsDisable": false,
								"draft":           false,
								"promoted":        false,
								"position":        0,
								"voteSum":         0,
								"voteCount":       0,
								"createdAt":       time.Date(2017, 12, 27, 3, 14, 16, 0, time.UTC),
								"updatedAt":       time.Date(2017, 12, 27, 3, 14, 48, 0, time.UTC),
								"sourceLocale":    "zh-tw",
								"outdated":        false,
								"outdatedLocales": []string{},
								"editedAt":        time.Date(2017, 12, 27, 3, 14, 16, 0, time.UTC),
								"labelNames":      []string{},
								"countryCode":     "tw",
								"url":             "https://honestbeehelp-tw.zendesk.com/api/v2/help_center/en-us/articles/115015959188-What-can-I-do-when-my-cart-is-locked-.json",
								"htmlUrl":         "https://help.honestbee.tw/hc/en-us/articles/115015959188-What-can-I-do-when-my-cart-is-locked-",
								"name":            "What can I do when my cart is locked?",
								"title":           "What can I do when my cart is locked?",
								"body":            "<p>When there is an error completing your checkout,your cart will be temporarily locked to prevent further changes to your order. To unlock your cart,click ‘Yes,unlock my cart,’ when prompted.</p>",
								"locale":          "en-us",
								"categoryConnection": map[string]interface{}{
									"id":      "115002432448",
									"name":    "My Account",
									"keyName": "myAccount",
								},
								"sectionConnection": map[string]interface{}{
									"id":   "115004118448",
									"name": "I need help with my account",
								},
							},
							map[string]interface{}{
								"id":              "115015885507",
								"authorId":        "24400224208",
								"commentsDisable": false,
								"draft":           false,
								"promoted":        false,
								"position":        0,
								"voteSum":         0,
								"voteCount":       0,
								"createdAt":       time.Date(2017, 12, 27, 3, 11, 30, 0, time.UTC),
								"updatedAt":       time.Date(2018, 1, 10, 6, 7, 41, 0, time.UTC),
								"sourceLocale":    "zh-tw",
								"outdated":        false,
								"outdatedLocales": []string{},
								"editedAt":        time.Date(2018, 1, 10, 6, 7, 41, 0, time.UTC),
								"labelNames":      []string{},
								"countryCode":     "tw",
								"url":             "https://honestbeehelp-tw.zendesk.com/api/v2/help_center/en-us/articles/115015885507-How-can-I-change-my-email-address-.json",
								"htmlUrl":         "https://help.honestbee.tw/hc/en-us/articles/115015885507-How-can-I-change-my-email-address-",
								"name":            "How can I change my email address?",
								"title":           "How can I change my email address?",
								"body":            `<p>Currently,it’s not possible to change your email address. To register with a different email address,please create a new account.</p>`,
								"locale":          "en-us",
								"categoryConnection": map[string]interface{}{
									"id":      "115002432448",
									"name":    "My Account",
									"keyName": "myAccount",
								},
								"sectionConnection": map[string]interface{}{
									"id":   "115004118448",
									"name": "I need help with my account",
								},
							},
							map[string]interface{}{
								"id":              "115015959148",
								"authorId":        "24400224208",
								"commentsDisable": false,
								"draft":           false,
								"promoted":        false,
								"position":        0,
								"voteSum":         -1,
								"voteCount":       1,
								"createdAt":       time.Date(2017, 12, 27, 2, 59, 48, 0, time.UTC),
								"updatedAt":       time.Date(2018, 3, 2, 5, 59, 58, 0, time.UTC),
								"sourceLocale":    "zh-tw",
								"outdated":        false,
								"outdatedLocales": []string{},
								"editedAt":        time.Date(2017, 12, 27, 3, 6, 54, 0, time.UTC),
								"labelNames":      []string{},
								"countryCode":     "tw",
								"url":             "https://honestbeehelp-tw.zendesk.com/api/v2/help_center/en-us/articles/115015959148-Help-I-ve-forgotten-my-password-.json",
								"htmlUrl":         "https://help.honestbee.tw/hc/en-us/articles/115015959148-Help-I-ve-forgotten-my-password-",
								"name":            "Help! I’ve forgotten my password.",
								"title":           "Help! I’ve forgotten my password.",
								"body":            "<p>Click on the Forgot Password link on the Login page and enter your registered email address. We’ll send you an email to reset your password. Occasionally emails end up in the junk/spam folder. Take a look there.</p>",
								"locale":          "en-us",
								"categoryConnection": map[string]interface{}{
									"id":      "115002432448",
									"name":    "My Account",
									"keyName": "myAccount",
								},
								"sectionConnection": map[string]interface{}{
									"id":   "115004118448",
									"name": "I need help with my account",
								},
							},
							map[string]interface{}{
								"id":              "115015959168",
								"authorId":        "24400224208",
								"commentsDisable": false,
								"draft":           false,
								"promoted":        false,
								"position":        0,
								"voteSum":         -2,
								"voteCount":       -2,
								"createdAt":       time.Date(2017, 12, 27, 3, 8, 5, 0, time.UTC),
								"updatedAt":       time.Date(2018, 3, 6, 12, 39, 30, 0, time.UTC),
								"sourceLocale":    "zh-tw",
								"outdated":        false,
								"outdatedLocales": []string{},
								"editedAt":        time.Date(2018, 1, 10, 6, 7, 41, 0, time.UTC),
								"labelNames":      []string{},
								"countryCode":     "tw",
								"url":             "https://honestbeehelp-tw.zendesk.com/api/v2/help_center/en-us/articles/115015959168-How-can-I-change-my-credit-card-password-or-delivery-address-.json",
								"htmlUrl":         "https://help.honestbee.tw/hc/en-us/articles/115015959168-How-can-I-change-my-credit-card-password-or-delivery-address-",
								"name":            "How can I change my credit card,password or delivery address?",
								"title":           "How can I change my credit card,password or delivery address?",
								"body":            `<p>Log in to your account and go to your profile icon at the top right corner. Select Settings from the dropdown menu to edit your details.</p>`,
								"locale":          "en-us",
								"categoryConnection": map[string]interface{}{
									"id":      "115002432448",
									"name":    "My Account",
									"keyName": "myAccount",
								},
								"sectionConnection": map[string]interface{}{
									"id":   "115004118448",
									"name": "I need help with my account",
								},
							},
						},
					},
				},
			},
		},
		{
			description: "testing sort order desc with order by updated at case",
			body: map[string]interface{}{
				"query": `
				{
					allArticles(countryCode: TW, locale: EN_US, sortBy: UPDATED_AT, sortOrder: DESC) {
						page
						perPage
						pageCount
						count
						articles {
							id
							authorId
							commentsDisable
							draft
							promoted
							position
							voteSum
							voteCount
							createdAt
							updatedAt
							sourceLocale
							outdated
							outdatedLocales
							editedAt
							labelNames
							countryCode
							url
							htmlUrl
							name
							title
							body
							locale
							categoryConnection {
								id
								name
								keyName
							}
							sectionConnection {
								id
								name
							}
						}
					}
				}
				`,
			},
			expectBody: map[string]interface{}{
				"data": map[string]interface{}{
					"allArticles": map[string]interface{}{
						"page":      1,
						"perPage":   30,
						"pageCount": 1,
						"count":     5,
						"articles": []interface{}{
							map[string]interface{}{
								"id":              "115015959168",
								"authorId":        "24400224208",
								"commentsDisable": false,
								"draft":           false,
								"promoted":        false,
								"position":        0,
								"voteSum":         -2,
								"voteCount":       -2,
								"createdAt":       time.Date(2017, 12, 27, 3, 8, 5, 0, time.UTC),
								"updatedAt":       time.Date(2018, 3, 6, 12, 39, 30, 0, time.UTC),
								"sourceLocale":    "zh-tw",
								"outdated":        false,
								"outdatedLocales": []string{},
								"editedAt":        time.Date(2018, 1, 10, 6, 7, 41, 0, time.UTC),
								"labelNames":      []string{},
								"countryCode":     "tw",
								"url":             "https://honestbeehelp-tw.zendesk.com/api/v2/help_center/en-us/articles/115015959168-How-can-I-change-my-credit-card-password-or-delivery-address-.json",
								"htmlUrl":         "https://help.honestbee.tw/hc/en-us/articles/115015959168-How-can-I-change-my-credit-card-password-or-delivery-address-",
								"name":            "How can I change my credit card,password or delivery address?",
								"title":           "How can I change my credit card,password or delivery address?",
								"body":            `<p>Log in to your account and go to your profile icon at the top right corner. Select Settings from the dropdown menu to edit your details.</p>`,
								"locale":          "en-us",
								"categoryConnection": map[string]interface{}{
									"id":      "115002432448",
									"name":    "My Account",
									"keyName": "myAccount",
								},
								"sectionConnection": map[string]interface{}{
									"id":   "115004118448",
									"name": "I need help with my account",
								},
							},
							map[string]interface{}{
								"id":              "115015959148",
								"authorId":        "24400224208",
								"commentsDisable": false,
								"draft":           false,
								"promoted":        false,
								"position":        0,
								"voteSum":         -1,
								"voteCount":       1,
								"createdAt":       time.Date(2017, 12, 27, 2, 59, 48, 0, time.UTC),
								"updatedAt":       time.Date(2018, 3, 2, 5, 59, 58, 0, time.UTC),
								"sourceLocale":    "zh-tw",
								"outdated":        false,
								"outdatedLocales": []string{},
								"editedAt":        time.Date(2017, 12, 27, 3, 6, 54, 0, time.UTC),
								"labelNames":      []string{},
								"countryCode":     "tw",
								"url":             "https://honestbeehelp-tw.zendesk.com/api/v2/help_center/en-us/articles/115015959148-Help-I-ve-forgotten-my-password-.json",
								"htmlUrl":         "https://help.honestbee.tw/hc/en-us/articles/115015959148-Help-I-ve-forgotten-my-password-",
								"name":            "Help! I’ve forgotten my password.",
								"title":           "Help! I’ve forgotten my password.",
								"body":            "<p>Click on the Forgot Password link on the Login page and enter your registered email address. We’ll send you an email to reset your password. Occasionally emails end up in the junk/spam folder. Take a look there.</p>",
								"locale":          "en-us",
								"categoryConnection": map[string]interface{}{
									"id":      "115002432448",
									"name":    "My Account",
									"keyName": "myAccount",
								},
								"sectionConnection": map[string]interface{}{
									"id":   "115004118448",
									"name": "I need help with my account",
								},
							},
							map[string]interface{}{
								"id":              "115015885507",
								"authorId":        "24400224208",
								"commentsDisable": false,
								"draft":           false,
								"promoted":        false,
								"position":        0,
								"voteSum":         0,
								"voteCount":       0,
								"createdAt":       time.Date(2017, 12, 27, 3, 11, 30, 0, time.UTC),
								"updatedAt":       time.Date(2018, 1, 10, 6, 7, 41, 0, time.UTC),
								"sourceLocale":    "zh-tw",
								"outdated":        false,
								"outdatedLocales": []string{},
								"editedAt":        time.Date(2018, 1, 10, 6, 7, 41, 0, time.UTC),
								"labelNames":      []string{},
								"countryCode":     "tw",
								"url":             "https://honestbeehelp-tw.zendesk.com/api/v2/help_center/en-us/articles/115015885507-How-can-I-change-my-email-address-.json",
								"htmlUrl":         "https://help.honestbee.tw/hc/en-us/articles/115015885507-How-can-I-change-my-email-address-",
								"name":            "How can I change my email address?",
								"title":           "How can I change my email address?",
								"body":            `<p>Currently,it’s not possible to change your email address. To register with a different email address,please create a new account.</p>`,
								"locale":          "en-us",
								"categoryConnection": map[string]interface{}{
									"id":      "115002432448",
									"name":    "My Account",
									"keyName": "myAccount",
								},
								"sectionConnection": map[string]interface{}{
									"id":   "115004118448",
									"name": "I need help with my account",
								},
							},
							map[string]interface{}{
								"id":              "115015959188",
								"authorId":        "24400224208",
								"commentsDisable": false,
								"draft":           false,
								"promoted":        false,
								"position":        0,
								"voteSum":         0,
								"voteCount":       0,
								"createdAt":       time.Date(2017, 12, 27, 3, 14, 16, 0, time.UTC),
								"updatedAt":       time.Date(2017, 12, 27, 3, 14, 48, 0, time.UTC),
								"sourceLocale":    "zh-tw",
								"outdated":        false,
								"outdatedLocales": []string{},
								"editedAt":        time.Date(2017, 12, 27, 3, 14, 16, 0, time.UTC),
								"labelNames":      []string{},
								"countryCode":     "tw",
								"url":             "https://honestbeehelp-tw.zendesk.com/api/v2/help_center/en-us/articles/115015959188-What-can-I-do-when-my-cart-is-locked-.json",
								"htmlUrl":         "https://help.honestbee.tw/hc/en-us/articles/115015959188-What-can-I-do-when-my-cart-is-locked-",
								"name":            "What can I do when my cart is locked?",
								"title":           "What can I do when my cart is locked?",
								"body":            "<p>When there is an error completing your checkout,your cart will be temporarily locked to prevent further changes to your order. To unlock your cart,click ‘Yes,unlock my cart,’ when prompted.</p>",
								"locale":          "en-us",
								"categoryConnection": map[string]interface{}{
									"id":      "115002432448",
									"name":    "My Account",
									"keyName": "myAccount",
								},
								"sectionConnection": map[string]interface{}{
									"id":   "115004118448",
									"name": "I need help with my account",
								},
							},
							map[string]interface{}{
								"id":              "115015885547",
								"authorId":        "24400224208",
								"commentsDisable": false,
								"draft":           false,
								"promoted":        false,
								"position":        0,
								"voteSum":         0,
								"voteCount":       0,
								"createdAt":       time.Date(2017, 12, 27, 3, 12, 58, 0, time.UTC),
								"updatedAt":       time.Date(2017, 12, 27, 3, 13, 42, 0, time.UTC),
								"sourceLocale":    "zh-tw",
								"outdated":        false,
								"outdatedLocales": []string{},
								"editedAt":        time.Date(2017, 12, 27, 3, 12, 58, 0, time.UTC),
								"labelNames":      []string{},
								"countryCode":     "tw",
								"url":             "https://honestbeehelp-tw.zendesk.com/api/v2/help_center/en-us/articles/115015885547-What-is-honestbee-s-information-security-policy-.json",
								"htmlUrl":         "https://help.honestbee.tw/hc/en-us/articles/115015885547-What-is-honestbee-s-information-security-policy-",
								"name":            "What is honestbee’s information security policy?",
								"title":           "What is honestbee’s information security policy?",
								"body":            `<p>honestbee takes your privacy seriously and complies with all the relevant laws to ensure your details are kept secure. Read our <a href="https://www.honestbee.tw/privacy-policy">Privacy Policy</a> for more information.</p>`,
								"locale":          "en-us",
								"categoryConnection": map[string]interface{}{
									"id":      "115002432448",
									"name":    "My Account",
									"keyName": "myAccount",
								},
								"sectionConnection": map[string]interface{}{
									"id":   "115004118448",
									"name": "I need help with my account",
								},
							},
						},
					},
				},
			},
		},
		{
			description: "testing per page 2 case",
			body: map[string]interface{}{
				"query": `
				{
					allArticles(countryCode: TW, locale: EN_US, perPage: 2) {
						page
						perPage
						pageCount
						count
						articles {
							id
							authorId
							commentsDisable
							draft
							promoted
							position
							voteSum
							voteCount
							createdAt
							updatedAt
							sourceLocale
							outdated
							outdatedLocales
							editedAt
							labelNames
							countryCode
							url
							htmlUrl
							name
							title
							body
							locale
							categoryConnection {
								id
								name
								keyName
							}
							sectionConnection {
								id
								name
							}
						}
					}
				}
				`,
			},
			expectBody: map[string]interface{}{
				"data": map[string]interface{}{
					"allArticles": map[string]interface{}{
						"page":      1,
						"perPage":   2,
						"pageCount": 3,
						"count":     5,
						"articles": []interface{}{
							map[string]interface{}{
								"id":              "115015959188",
								"authorId":        "24400224208",
								"commentsDisable": false,
								"draft":           false,
								"promoted":        false,
								"position":        0,
								"voteSum":         0,
								"voteCount":       0,
								"createdAt":       time.Date(2017, 12, 27, 3, 14, 16, 0, time.UTC),
								"updatedAt":       time.Date(2017, 12, 27, 3, 14, 48, 0, time.UTC),
								"sourceLocale":    "zh-tw",
								"outdated":        false,
								"outdatedLocales": []string{},
								"editedAt":        time.Date(2017, 12, 27, 3, 14, 16, 0, time.UTC),
								"labelNames":      []string{},
								"countryCode":     "tw",
								"url":             "https://honestbeehelp-tw.zendesk.com/api/v2/help_center/en-us/articles/115015959188-What-can-I-do-when-my-cart-is-locked-.json",
								"htmlUrl":         "https://help.honestbee.tw/hc/en-us/articles/115015959188-What-can-I-do-when-my-cart-is-locked-",
								"name":            "What can I do when my cart is locked?",
								"title":           "What can I do when my cart is locked?",
								"body":            "<p>When there is an error completing your checkout,your cart will be temporarily locked to prevent further changes to your order. To unlock your cart,click ‘Yes,unlock my cart,’ when prompted.</p>",
								"locale":          "en-us",
								"categoryConnection": map[string]interface{}{
									"id":      "115002432448",
									"name":    "My Account",
									"keyName": "myAccount",
								},
								"sectionConnection": map[string]interface{}{
									"id":   "115004118448",
									"name": "I need help with my account",
								},
							},
							map[string]interface{}{
								"id":              "115015885547",
								"authorId":        "24400224208",
								"commentsDisable": false,
								"draft":           false,
								"promoted":        false,
								"position":        0,
								"voteSum":         0,
								"voteCount":       0,
								"createdAt":       time.Date(2017, 12, 27, 3, 12, 58, 0, time.UTC),
								"updatedAt":       time.Date(2017, 12, 27, 3, 13, 42, 0, time.UTC),
								"sourceLocale":    "zh-tw",
								"outdated":        false,
								"outdatedLocales": []string{},
								"editedAt":        time.Date(2017, 12, 27, 3, 12, 58, 0, time.UTC),
								"labelNames":      []string{},
								"countryCode":     "tw",
								"url":             "https://honestbeehelp-tw.zendesk.com/api/v2/help_center/en-us/articles/115015885547-What-is-honestbee-s-information-security-policy-.json",
								"htmlUrl":         "https://help.honestbee.tw/hc/en-us/articles/115015885547-What-is-honestbee-s-information-security-policy-",
								"name":            "What is honestbee’s information security policy?",
								"title":           "What is honestbee’s information security policy?",
								"body":            `<p>honestbee takes your privacy seriously and complies with all the relevant laws to ensure your details are kept secure. Read our <a href="https://www.honestbee.tw/privacy-policy">Privacy Policy</a> for more information.</p>`,
								"locale":          "en-us",
								"categoryConnection": map[string]interface{}{
									"id":      "115002432448",
									"name":    "My Account",
									"keyName": "myAccount",
								},
								"sectionConnection": map[string]interface{}{
									"id":   "115004118448",
									"name": "I need help with my account",
								},
							},
						},
					},
				},
			},
		},
		{
			description: "testing per page 2 page 2 case",
			body: map[string]interface{}{
				"query": `
				{
					allArticles(countryCode: TW, locale: EN_US, perPage: 2, page: 2) {
						page
						perPage
						pageCount
						count
						articles {
							id
							authorId
							commentsDisable
							draft
							promoted
							position
							voteSum
							voteCount
							createdAt
							updatedAt
							sourceLocale
							outdated
							outdatedLocales
							editedAt
							labelNames
							countryCode
							url
							htmlUrl
							name
							title
							body
							locale
							categoryConnection {
								id
								name
								keyName
							}
							sectionConnection {
								id
								name
							}
						}
					}
				}
				`,
			},
			expectBody: map[string]interface{}{
				"data": map[string]interface{}{
					"allArticles": map[string]interface{}{
						"page":      2,
						"perPage":   2,
						"pageCount": 3,
						"count":     5,
						"articles": []interface{}{
							map[string]interface{}{
								"id":              "115015885507",
								"authorId":        "24400224208",
								"commentsDisable": false,
								"draft":           false,
								"promoted":        false,
								"position":        0,
								"voteSum":         0,
								"voteCount":       0,
								"createdAt":       time.Date(2017, 12, 27, 3, 11, 30, 0, time.UTC),
								"updatedAt":       time.Date(2018, 1, 10, 6, 7, 41, 0, time.UTC),
								"sourceLocale":    "zh-tw",
								"outdated":        false,
								"outdatedLocales": []string{},
								"editedAt":        time.Date(2018, 1, 10, 6, 7, 41, 0, time.UTC),
								"labelNames":      []string{},
								"countryCode":     "tw",
								"url":             "https://honestbeehelp-tw.zendesk.com/api/v2/help_center/en-us/articles/115015885507-How-can-I-change-my-email-address-.json",
								"htmlUrl":         "https://help.honestbee.tw/hc/en-us/articles/115015885507-How-can-I-change-my-email-address-",
								"name":            "How can I change my email address?",
								"title":           "How can I change my email address?",
								"body":            `<p>Currently,it’s not possible to change your email address. To register with a different email address,please create a new account.</p>`,
								"locale":          "en-us",
								"categoryConnection": map[string]interface{}{
									"id":      "115002432448",
									"name":    "My Account",
									"keyName": "myAccount",
								},
								"sectionConnection": map[string]interface{}{
									"id":   "115004118448",
									"name": "I need help with my account",
								},
							},
							map[string]interface{}{
								"id":              "115015959168",
								"authorId":        "24400224208",
								"commentsDisable": false,
								"draft":           false,
								"promoted":        false,
								"position":        0,
								"voteSum":         -2,
								"voteCount":       -2,
								"createdAt":       time.Date(2017, 12, 27, 3, 8, 5, 0, time.UTC),
								"updatedAt":       time.Date(2018, 3, 6, 12, 39, 30, 0, time.UTC),
								"sourceLocale":    "zh-tw",
								"outdated":        false,
								"outdatedLocales": []string{},
								"editedAt":        time.Date(2018, 1, 10, 6, 7, 41, 0, time.UTC),
								"labelNames":      []string{},
								"countryCode":     "tw",
								"url":             "https://honestbeehelp-tw.zendesk.com/api/v2/help_center/en-us/articles/115015959168-How-can-I-change-my-credit-card-password-or-delivery-address-.json",
								"htmlUrl":         "https://help.honestbee.tw/hc/en-us/articles/115015959168-How-can-I-change-my-credit-card-password-or-delivery-address-",
								"name":            "How can I change my credit card,password or delivery address?",
								"title":           "How can I change my credit card,password or delivery address?",
								"body":            `<p>Log in to your account and go to your profile icon at the top right corner. Select Settings from the dropdown menu to edit your details.</p>`,
								"locale":          "en-us",
								"categoryConnection": map[string]interface{}{
									"id":      "115002432448",
									"name":    "My Account",
									"keyName": "myAccount",
								},
								"sectionConnection": map[string]interface{}{
									"id":   "115004118448",
									"name": "I need help with my account",
								},
							},
						},
					},
				},
			},
		},
		{
			description: "testing per page 2 page 3 case",
			body: map[string]interface{}{
				"query": `
				{
					allArticles(countryCode: TW, locale: EN_US, perPage: 2, page: 3) {
						page
						perPage
						pageCount
						count
						articles {
							id
							authorId
							commentsDisable
							draft
							promoted
							position
							voteSum
							voteCount
							createdAt
							updatedAt
							sourceLocale
							outdated
							outdatedLocales
							editedAt
							labelNames
							countryCode
							url
							htmlUrl
							name
							title
							body
							locale
							categoryConnection {
								id
								name
								keyName
							}
							sectionConnection {
								id
								name
							}
						}
					}
				}
				`,
			},
			expectBody: map[string]interface{}{
				"data": map[string]interface{}{
					"allArticles": map[string]interface{}{
						"page":      3,
						"perPage":   2,
						"pageCount": 3,
						"count":     5,
						"articles": []interface{}{
							map[string]interface{}{
								"id":              "115015959148",
								"authorId":        "24400224208",
								"commentsDisable": false,
								"draft":           false,
								"promoted":        false,
								"position":        0,
								"voteSum":         -1,
								"voteCount":       1,
								"createdAt":       time.Date(2017, 12, 27, 2, 59, 48, 0, time.UTC),
								"updatedAt":       time.Date(2018, 3, 2, 5, 59, 58, 0, time.UTC),
								"sourceLocale":    "zh-tw",
								"outdated":        false,
								"outdatedLocales": []string{},
								"editedAt":        time.Date(2017, 12, 27, 3, 6, 54, 0, time.UTC),
								"labelNames":      []string{},
								"countryCode":     "tw",
								"url":             "https://honestbeehelp-tw.zendesk.com/api/v2/help_center/en-us/articles/115015959148-Help-I-ve-forgotten-my-password-.json",
								"htmlUrl":         "https://help.honestbee.tw/hc/en-us/articles/115015959148-Help-I-ve-forgotten-my-password-",
								"name":            "Help! I’ve forgotten my password.",
								"title":           "Help! I’ve forgotten my password.",
								"body":            "<p>Click on the Forgot Password link on the Login page and enter your registered email address. We’ll send you an email to reset your password. Occasionally emails end up in the junk/spam folder. Take a look there.</p>",
								"locale":          "en-us",
								"categoryConnection": map[string]interface{}{
									"id":      "115002432448",
									"name":    "My Account",
									"keyName": "myAccount",
								},
								"sectionConnection": map[string]interface{}{
									"id":   "115004118448",
									"name": "I need help with my account",
								},
							},
						},
					},
				},
			},
		},
		{
			description: "testing not exist country code case",
			body: map[string]interface{}{
				"query": `
				{
					allArticles(countryCode: not_exist_country_code) {
						page
						perPage
						pageCount
						count
					}
				}
				`,
			},
			expectBody: map[string]interface{}{
				"errors": []interface{}{
					map[string]interface{}{
						"message": "Argument \"countryCode\" has invalid value not_exist_country_code.\nExpected type \"CountryCode\", found not_exist_country_code.",
						"locations": []interface{}{
							map[string]interface{}{
								"line":   3,
								"column": 31,
							},
						},
					},
				},
			},
		},
		{
			description: "testing not exist locale case",
			body: map[string]interface{}{
				"query": `
				{
					allArticles(locale: not_exist_locale) {
						page
						perPage
						pageCount
						count
					}
				}
				`,
			},
			expectBody: map[string]interface{}{
				"errors": []interface{}{
					map[string]interface{}{
						"message": "Argument \"locale\" has invalid value not_exist_locale.\nExpected type \"Locale\", found not_exist_locale.",
						"locations": []interface{}{
							map[string]interface{}{
								"line":   3,
								"column": 26,
							},
						},
					},
				},
			},
		},
		{
			description: "testing sort by not correct case",
			body: map[string]interface{}{
				"query": `
				{
					allArticles(sortBy: unknown_sort_by) {
						page
						perPage
						pageCount
						count
					}
				}
				`,
			},
			expectBody: map[string]interface{}{
				"errors": []interface{}{
					map[string]interface{}{
						"message": "Argument \"sortBy\" has invalid value unknown_sort_by.\nExpected type \"SortBy\", found unknown_sort_by.",
						"locations": []interface{}{
							map[string]interface{}{
								"line":   3,
								"column": 26,
							},
						},
					},
				},
			},
		},
		{
			description: "testing sort order not correct case",
			body: map[string]interface{}{
				"query": `
				{
					allArticles(sortOrder: unknown_sort_order) {
						page
						perPage
						pageCount
						count
					}
				}
				`,
			},
			expectBody: map[string]interface{}{
				"errors": []interface{}{
					map[string]interface{}{
						"message": "Argument \"sortOrder\" has invalid value unknown_sort_order.\nExpected type \"SortOrder\", found unknown_sort_order.",
						"locations": []interface{}{
							map[string]interface{}{
								"line":   3,
								"column": 29,
							},
						},
					},
				},
			},
		},
	}

	for _, tt := range testCases {
		t.Run(tt.description, func(t *testing.T) {
			// Send requests.
			b, err := json.Marshal(tt.body)
			if err != nil {
				t.Fatalf("[%s] json marshal failed:%v", tt.description, err)
			}
			resp, err := ts.Client().Post(ts.URL+"/graphql", "application/json", ioutil.NopCloser(bytes.NewReader(b)))
			if err != nil {
				t.Fatalf("[%s] http client get failed:%v", tt.description, err)
			}
			defer resp.Body.Close()

			// Compare HTTP status code.
			if http.StatusOK != resp.StatusCode {
				t.Errorf("[%s] http status expect:%v != actual:%v", tt.description, http.StatusOK, resp.StatusCode)
			}

			// Compare HTTP body.
			actual := make(map[string]interface{})
			if err = json.NewDecoder(resp.Body).Decode(&actual); err != nil {
				t.Fatalf("[%s] json decoding failed:%v", tt.description, err)
			}
			// Converts integer to the same type.
			expectData, err := json.Marshal(tt.expectBody)
			if err != nil {
				t.Fatalf("[%s] json marshal failed:%v", tt.description, err)
			}
			expect := make(map[string]interface{})
			if err = json.Unmarshal(expectData, &expect); err != nil {
				t.Fatalf("[%s] json unmarshal failed:%v", tt.description, err)
			}
			// Compares and prints difference.
			if diff := deep.Equal(expect, actual); diff != nil {
				t.Errorf("[%s] %v", tt.description, diff)
			}
		})
	}
}

func TestHandlersGraphQLQueryTopArticles(t *testing.T) {
	ts := newTserver()
	defer ts.closeAll()
	testCases := []struct {
		description string
		body        map[string]interface{}
		query       string
		expectBody  map[string]interface{}
	}{
		{
			description: "testing normal top 3 TW + EN_US locale case",
			body: map[string]interface{}{
				"query": `
				{
					topArticles(topN: 3, countryCode: TW, locale: EN_US) {
						id
						authorId
						commentsDisable
						draft
						promoted
						position
						voteSum
						voteCount
						createdAt
						updatedAt
						sourceLocale
						outdated
						outdatedLocales
						editedAt
						labelNames
						countryCode
						url
						htmlUrl
						name
						title
						body
						locale
						categoryConnection {
							id
							name
							keyName
						}
						sectionConnection {
							id
							name
						}
					}
				}
				`,
			},
			expectBody: map[string]interface{}{
				"data": map[string]interface{}{
					"topArticles": []interface{}{
						map[string]interface{}{
							"id":              "115015959188",
							"authorId":        "24400224208",
							"commentsDisable": false,
							"draft":           false,
							"promoted":        false,
							"position":        0,
							"voteSum":         0,
							"voteCount":       0,
							"createdAt":       time.Date(2017, 12, 27, 3, 14, 16, 0, time.UTC),
							"updatedAt":       time.Date(2017, 12, 27, 3, 14, 48, 0, time.UTC),
							"sourceLocale":    "zh-tw",
							"outdated":        false,
							"outdatedLocales": []string{},
							"editedAt":        time.Date(2017, 12, 27, 3, 14, 16, 0, time.UTC),
							"labelNames":      []string{},
							"countryCode":     "tw",
							"url":             "https://honestbeehelp-tw.zendesk.com/api/v2/help_center/en-us/articles/115015959188-What-can-I-do-when-my-cart-is-locked-.json",
							"htmlUrl":         "https://help.honestbee.tw/hc/en-us/articles/115015959188-What-can-I-do-when-my-cart-is-locked-",
							"name":            "What can I do when my cart is locked?",
							"title":           "What can I do when my cart is locked?",
							"body":            "<p>When there is an error completing your checkout,your cart will be temporarily locked to prevent further changes to your order. To unlock your cart,click ‘Yes,unlock my cart,’ when prompted.</p>",
							"locale":          "en-us",
							"categoryConnection": map[string]interface{}{
								"id":      "115002432448",
								"name":    "My Account",
								"keyName": "myAccount",
							},
							"sectionConnection": map[string]interface{}{
								"id":   "115004118448",
								"name": "I need help with my account",
							},
						},
						map[string]interface{}{
							"id":              "115015959148",
							"authorId":        "24400224208",
							"commentsDisable": false,
							"draft":           false,
							"promoted":        false,
							"position":        0,
							"voteSum":         -1,
							"voteCount":       1,
							"createdAt":       time.Date(2017, 12, 27, 2, 59, 48, 0, time.UTC),
							"updatedAt":       time.Date(2018, 3, 2, 5, 59, 58, 0, time.UTC),
							"sourceLocale":    "zh-tw",
							"outdated":        false,
							"outdatedLocales": []string{},
							"editedAt":        time.Date(2017, 12, 27, 3, 6, 54, 0, time.UTC),
							"labelNames":      []string{},
							"countryCode":     "tw",
							"url":             "https://honestbeehelp-tw.zendesk.com/api/v2/help_center/en-us/articles/115015959148-Help-I-ve-forgotten-my-password-.json",
							"htmlUrl":         "https://help.honestbee.tw/hc/en-us/articles/115015959148-Help-I-ve-forgotten-my-password-",
							"name":            "Help! I’ve forgotten my password.",
							"title":           "Help! I’ve forgotten my password.",
							"body":            "<p>Click on the Forgot Password link on the Login page and enter your registered email address. We’ll send you an email to reset your password. Occasionally emails end up in the junk/spam folder. Take a look there.</p>",
							"locale":          "en-us",
							"categoryConnection": map[string]interface{}{
								"id":      "115002432448",
								"name":    "My Account",
								"keyName": "myAccount",
							},
							"sectionConnection": map[string]interface{}{
								"id":   "115004118448",
								"name": "I need help with my account",
							},
						},
						map[string]interface{}{
							"id":              "115015885547",
							"authorId":        "24400224208",
							"commentsDisable": false,
							"draft":           false,
							"promoted":        false,
							"position":        0,
							"voteSum":         0,
							"voteCount":       0,
							"createdAt":       time.Date(2017, 12, 27, 3, 12, 58, 0, time.UTC),
							"updatedAt":       time.Date(2017, 12, 27, 3, 13, 42, 0, time.UTC),
							"sourceLocale":    "zh-tw",
							"outdated":        false,
							"outdatedLocales": []string{},
							"editedAt":        time.Date(2017, 12, 27, 3, 12, 58, 0, time.UTC),
							"labelNames":      []string{},
							"countryCode":     "tw",
							"url":             "https://honestbeehelp-tw.zendesk.com/api/v2/help_center/en-us/articles/115015885547-What-is-honestbee-s-information-security-policy-.json",
							"htmlUrl":         "https://help.honestbee.tw/hc/en-us/articles/115015885547-What-is-honestbee-s-information-security-policy-",
							"name":            "What is honestbee’s information security policy?",
							"title":           "What is honestbee’s information security policy?",
							"body":            `<p>honestbee takes your privacy seriously and complies with all the relevant laws to ensure your details are kept secure. Read our <a href="https://www.honestbee.tw/privacy-policy">Privacy Policy</a> for more information.</p>`,
							"locale":          "en-us",
							"categoryConnection": map[string]interface{}{
								"id":      "115002432448",
								"name":    "My Account",
								"keyName": "myAccount",
							},
							"sectionConnection": map[string]interface{}{
								"id":   "115004118448",
								"name": "I need help with my account",
							},
						},
					},
				},
			},
		},
		{
			description: "testing normal top 2 TW + EN_US locale case",
			body: map[string]interface{}{
				"query": `
				{
					topArticles(topN: 2, countryCode: TW, locale: EN_US) {
						id
						authorId
						commentsDisable
						draft
						promoted
						position
						voteSum
						voteCount
						createdAt
						updatedAt
						sourceLocale
						outdated
						outdatedLocales
						editedAt
						labelNames
						countryCode
						url
						htmlUrl
						name
						title
						body
						locale
						categoryConnection {
							id
							name
							keyName
						}
						sectionConnection {
							id
							name
						}
					}
				}
				`,
			},
			expectBody: map[string]interface{}{
				"data": map[string]interface{}{
					"topArticles": []interface{}{
						map[string]interface{}{
							"id":              "115015959188",
							"authorId":        "24400224208",
							"commentsDisable": false,
							"draft":           false,
							"promoted":        false,
							"position":        0,
							"voteSum":         0,
							"voteCount":       0,
							"createdAt":       time.Date(2017, 12, 27, 3, 14, 16, 0, time.UTC),
							"updatedAt":       time.Date(2017, 12, 27, 3, 14, 48, 0, time.UTC),
							"sourceLocale":    "zh-tw",
							"outdated":        false,
							"outdatedLocales": []string{},
							"editedAt":        time.Date(2017, 12, 27, 3, 14, 16, 0, time.UTC),
							"labelNames":      []string{},
							"countryCode":     "tw",
							"url":             "https://honestbeehelp-tw.zendesk.com/api/v2/help_center/en-us/articles/115015959188-What-can-I-do-when-my-cart-is-locked-.json",
							"htmlUrl":         "https://help.honestbee.tw/hc/en-us/articles/115015959188-What-can-I-do-when-my-cart-is-locked-",
							"name":            "What can I do when my cart is locked?",
							"title":           "What can I do when my cart is locked?",
							"body":            "<p>When there is an error completing your checkout,your cart will be temporarily locked to prevent further changes to your order. To unlock your cart,click ‘Yes,unlock my cart,’ when prompted.</p>",
							"locale":          "en-us",
							"categoryConnection": map[string]interface{}{
								"id":      "115002432448",
								"name":    "My Account",
								"keyName": "myAccount",
							},
							"sectionConnection": map[string]interface{}{
								"id":   "115004118448",
								"name": "I need help with my account",
							},
						},
						map[string]interface{}{
							"id":              "115015959148",
							"authorId":        "24400224208",
							"commentsDisable": false,
							"draft":           false,
							"promoted":        false,
							"position":        0,
							"voteSum":         -1,
							"voteCount":       1,
							"createdAt":       time.Date(2017, 12, 27, 2, 59, 48, 0, time.UTC),
							"updatedAt":       time.Date(2018, 3, 2, 5, 59, 58, 0, time.UTC),
							"sourceLocale":    "zh-tw",
							"outdated":        false,
							"outdatedLocales": []string{},
							"editedAt":        time.Date(2017, 12, 27, 3, 6, 54, 0, time.UTC),
							"labelNames":      []string{},
							"countryCode":     "tw",
							"url":             "https://honestbeehelp-tw.zendesk.com/api/v2/help_center/en-us/articles/115015959148-Help-I-ve-forgotten-my-password-.json",
							"htmlUrl":         "https://help.honestbee.tw/hc/en-us/articles/115015959148-Help-I-ve-forgotten-my-password-",
							"name":            "Help! I’ve forgotten my password.",
							"title":           "Help! I’ve forgotten my password.",
							"body":            "<p>Click on the Forgot Password link on the Login page and enter your registered email address. We’ll send you an email to reset your password. Occasionally emails end up in the junk/spam folder. Take a look there.</p>",
							"locale":          "en-us",
							"categoryConnection": map[string]interface{}{
								"id":      "115002432448",
								"name":    "My Account",
								"keyName": "myAccount",
							},
							"sectionConnection": map[string]interface{}{
								"id":   "115004118448",
								"name": "I need help with my account",
							},
						},
					},
				},
			},
		},
		{
			description: "testing normal top 3 TW + ZH_TW locale case",
			body: map[string]interface{}{
				"query": `
				{
					topArticles(topN: 3, countryCode: TW, locale: ZH_TW) {
						id
						authorId
						commentsDisable
						draft
						promoted
						position
						voteSum
						voteCount
						createdAt
						updatedAt
						sourceLocale
						outdated
						outdatedLocales
						editedAt
						labelNames
						countryCode
						url
						htmlUrl
						name
						title
						body
						locale
						categoryConnection {
							id
							name
							keyName
						}
						sectionConnection {
							id
							name
						}
					}
				}
				`,
			},
			expectBody: map[string]interface{}{
				"data": map[string]interface{}{
					"topArticles": []interface{}{
						map[string]interface{}{
							"id":              "115015959188",
							"authorId":        "24400224208",
							"commentsDisable": false,
							"draft":           false,
							"promoted":        false,
							"position":        0,
							"voteSum":         0,
							"voteCount":       0,
							"createdAt":       time.Date(2017, 12, 27, 3, 14, 16, 0, time.UTC),
							"updatedAt":       time.Date(2017, 12, 27, 3, 14, 48, 0, time.UTC),
							"sourceLocale":    "zh-tw",
							"outdated":        false,
							"outdatedLocales": []string{},
							"editedAt":        time.Date(2017, 12, 27, 3, 14, 16, 0, time.UTC),
							"labelNames":      []string{},
							"countryCode":     "tw",
							"url":             "https://honestbeehelp-tw.zendesk.com/api/v2/help_center/zh-tw/articles/115015959188-%E6%88%91%E7%9A%84%E8%B3%BC%E7%89%A9%E8%BB%8A%E9%8E%96%E4%BD%8F%E4%BA%86-%E8%A9%B2%E6%80%8E%E9%BA%BC%E8%BE%A6-.json",
							"htmlUrl":         "https://help.honestbee.tw/hc/zh-tw/articles/115015959188-%E6%88%91%E7%9A%84%E8%B3%BC%E7%89%A9%E8%BB%8A%E9%8E%96%E4%BD%8F%E4%BA%86-%E8%A9%B2%E6%80%8E%E9%BA%BC%E8%BE%A6-",
							"name":            "我的購物車鎖住了，該怎麼辦？",
							"title":           "我的購物車鎖住了，該怎麼辦？",
							"body":            "<p>當結帳出現錯誤時，您的購物車會暫時被鎖住，以避免您的訂單出現異動。要解鎖您的購物車，請在出現提示時，點選「是的，解鎖我的購物車」。</p>",
							"locale":          "zh-tw",
							"categoryConnection": map[string]interface{}{
								"id":      "115002432448",
								"name":    "我的帳號",
								"keyName": "myAccount",
							},
							"sectionConnection": map[string]interface{}{
								"id":   "115004118448",
								"name": "我需要帳號相關的協助",
							},
						},
						map[string]interface{}{
							"id":              "115015959148",
							"authorId":        "24400224208",
							"commentsDisable": false,
							"draft":           false,
							"promoted":        false,
							"position":        0,
							"voteSum":         -1,
							"voteCount":       1,
							"createdAt":       time.Date(2017, 12, 27, 2, 59, 48, 0, time.UTC),
							"updatedAt":       time.Date(2018, 3, 2, 5, 59, 58, 0, time.UTC),
							"sourceLocale":    "zh-tw",
							"outdated":        false,
							"outdatedLocales": []string{},
							"editedAt":        time.Date(2017, 12, 27, 3, 6, 54, 0, time.UTC),
							"labelNames":      []string{},
							"countryCode":     "tw",
							"url":             "https://honestbeehelp-tw.zendesk.com/api/v2/help_center/zh-tw/articles/115015959148-%E6%95%91%E5%91%BD-%E6%88%91%E5%BF%98%E8%A8%98%E5%AF%86%E7%A2%BC%E4%BA%86.json",
							"htmlUrl":         "https://help.honestbee.tw/hc/zh-tw/articles/115015959148-%E6%95%91%E5%91%BD-%E6%88%91%E5%BF%98%E8%A8%98%E5%AF%86%E7%A2%BC%E4%BA%86",
							"name":            "救命！我忘記密碼了",
							"title":           "救命！我忘記密碼了",
							"body":            "<p>請在登入頁面點選「忘記密碼」，並輸入您註冊時所使用的電子信箱。我們會寄給您一封信讓您重設密碼。有時候信件會跑到垃圾信件匣，請務必檢查看看。</p>",
							"locale":          "zh-tw",
							"categoryConnection": map[string]interface{}{
								"id":      "115002432448",
								"name":    "我的帳號",
								"keyName": "myAccount",
							},
							"sectionConnection": map[string]interface{}{
								"id":   "115004118448",
								"name": "我需要帳號相關的協助",
							},
						},
						map[string]interface{}{
							"id":              "115015885547",
							"authorId":        "24400224208",
							"commentsDisable": false,
							"draft":           false,
							"promoted":        false,
							"position":        0,
							"voteSum":         0,
							"voteCount":       0,
							"createdAt":       time.Date(2017, 12, 27, 3, 12, 58, 0, time.UTC),
							"updatedAt":       time.Date(2017, 12, 27, 3, 13, 42, 0, time.UTC),
							"sourceLocale":    "zh-tw",
							"outdated":        false,
							"outdatedLocales": []string{},
							"editedAt":        time.Date(2017, 12, 27, 3, 12, 58, 0, time.UTC),
							"labelNames":      []string{},
							"countryCode":     "tw",
							"url":             "https://honestbeehelp-tw.zendesk.com/api/v2/help_center/zh-tw/articles/115015885547-honestbee-%E7%9A%84%E5%AE%89%E5%85%A8%E6%94%BF%E7%AD%96%E7%82%BA%E4%BD%95-.json",
							"htmlUrl":         "https://help.honestbee.tw/hc/zh-tw/articles/115015885547-honestbee-%E7%9A%84%E5%AE%89%E5%85%A8%E6%94%BF%E7%AD%96%E7%82%BA%E4%BD%95-",
							"name":            "honestbee 的安全政策為何？",
							"title":           "honestbee 的安全政策為何？",
							"body":            `<p>honestbee 很重視您的隱私，並遵循所有相關法規以確保您的資訊安全。請閱讀我們的<a href="https://www.honestbee.tw/privacy-policy">隱私權政策</a>以了解更多資訊。</p>`,
							"locale":          "zh-tw",
							"categoryConnection": map[string]interface{}{
								"id":      "115002432448",
								"name":    "我的帳號",
								"keyName": "myAccount",
							},
							"sectionConnection": map[string]interface{}{
								"id":   "115004118448",
								"name": "我需要帳號相關的協助",
							},
						},
					},
				},
			},
		},
		{
			description: "testing negative top n input value case",
			body: map[string]interface{}{
				"query": `
				{
					topArticles(topN: -1, countryCode: TW, locale: EN_US) {
						id
						authorId
						commentsDisable
						draft
						promoted
						position
						voteSum
						voteCount
						createdAt
						updatedAt
						sourceLocale
						outdated
						outdatedLocales
						editedAt
						labelNames
						countryCode
						url
						htmlUrl
						name
						title
						body
						locale
						categoryConnection {
							id
							name
							keyName
						}
						sectionConnection {
							id
							name
						}
					}
				}
				`,
			},
			expectBody: map[string]interface{}{
				"data": map[string]interface{}{
					"topArticles": nil,
				},
				"errors": []interface{}{
					map[string]interface{}{
						"message": "You passed an invalid value for the attributes.",
						"path": []interface{}{
							"topArticles",
						},
					},
				},
			},
		},
		{
			description: "testing float top n input value case",
			body: map[string]interface{}{
				"query": `
				{
					topArticles(topN: 0.0, countryCode: TW, locale: EN_US) {
						id
						authorId
						commentsDisable
						draft
						promoted
						position
						voteSum
						voteCount
						createdAt
						updatedAt
						sourceLocale
						outdated
						outdatedLocales
						editedAt
						labelNames
						countryCode
						url
						htmlUrl
						name
						title
						body
						locale
						categoryConnection {
							id
							name
							keyName
						}
						sectionConnection {
							id
							name
						}
					}
				}
				`,
			},
			expectBody: map[string]interface{}{
				"errors": []interface{}{
					map[string]interface{}{
						"message": "Argument \"topN\" has invalid value 0.0.\nExpected type \"Int\", found 0.0.",
						"locations": []interface{}{
							map[string]interface{}{
								"line":   3,
								"column": 24,
							},
						},
					},
				},
			},
		},
		{
			description: "testing not exist country code case",
			body: map[string]interface{}{
				"query": `
				{
					topArticles(topN: 3, countryCode: not_exist_country_code) {
						id
						authorId
						commentsDisable
						draft
						promoted
						position
						voteSum
						voteCount
						createdAt
						updatedAt
						sourceLocale
						outdated
						outdatedLocales
						editedAt
						labelNames
						countryCode
						url
						htmlUrl
						name
						title
						body
						locale
						categoryConnection {
							id
							name
							keyName
						}
						sectionConnection {
							id
							name
						}
					}
				}
				`,
			},
			expectBody: map[string]interface{}{
				"errors": []interface{}{
					map[string]interface{}{
						"message": "Argument \"countryCode\" has invalid value not_exist_country_code.\nExpected type \"CountryCode\", found not_exist_country_code.",
						"locations": []interface{}{
							map[string]interface{}{
								"line":   3,
								"column": 40,
							},
						},
					},
				},
			},
		},
		{
			description: "testing not exist locale case",
			body: map[string]interface{}{
				"query": `
				{
					topArticles(topN: 3, locale: not_exist_locale) {
						id
						authorId
						commentsDisable
						draft
						promoted
						position
						voteSum
						voteCount
						createdAt
						updatedAt
						sourceLocale
						outdated
						outdatedLocales
						editedAt
						labelNames
						countryCode
						url
						htmlUrl
						name
						title
						body
						locale
						categoryConnection {
							id
							name
							keyName
						}
						sectionConnection {
							id
							name
						}
					}
				}
				`,
			},
			expectBody: map[string]interface{}{
				"errors": []interface{}{
					map[string]interface{}{
						"message": "Argument \"locale\" has invalid value not_exist_locale.\nExpected type \"Locale\", found not_exist_locale.",
						"locations": []interface{}{
							map[string]interface{}{
								"line":   3,
								"column": 35,
							},
						},
					},
				},
			},
		},
	}

	for _, tt := range testCases {
		t.Run(tt.description, func(t *testing.T) {
			// Send requests.
			b, err := json.Marshal(tt.body)
			if err != nil {
				t.Fatalf("[%s] json marshal failed:%v", tt.description, err)
			}
			resp, err := ts.Client().Post(ts.URL+"/graphql", "application/json", ioutil.NopCloser(bytes.NewReader(b)))
			if err != nil {
				t.Fatalf("[%s] http client get failed:%v", tt.description, err)
			}
			defer resp.Body.Close()

			// Compare HTTP status code.
			if http.StatusOK != resp.StatusCode {
				t.Errorf("[%s] http status expect:%v != actual:%v", tt.description, http.StatusOK, resp.StatusCode)
			}

			// Compare HTTP body.
			actual := make(map[string]interface{})
			if err = json.NewDecoder(resp.Body).Decode(&actual); err != nil {
				t.Fatalf("[%s] json decoding failed:%v", tt.description, err)
			}
			// Converts integer to the same type.
			expectData, err := json.Marshal(tt.expectBody)
			if err != nil {
				t.Fatalf("[%s] json marshal failed:%v", tt.description, err)
			}
			expect := make(map[string]interface{})
			if err = json.Unmarshal(expectData, &expect); err != nil {
				t.Fatalf("[%s] json unmarshal failed:%v", tt.description, err)
			}
			// Compares and prints difference.
			if diff := deep.Equal(expect, actual); diff != nil {
				t.Errorf("[%s] %v", tt.description, diff)
			}
		})
	}
}

func TestHandlersGraphQLQueryArticle(t *testing.T) {
	ts := newTserver()
	defer ts.closeAll()
	testCases := []struct {
		description string
		body        map[string]interface{}
		query       string
		expectBody  map[string]interface{}
	}{
		{
			description: "testing normal tw + en-us locale case",
			body: map[string]interface{}{
				"query": `
				{
					oneArticle(articleId: "115015959188", countryCode: TW, locale: EN_US) {
						id
						authorId
						commentsDisable
						draft
						promoted
						position
						voteSum
						voteCount
						createdAt
						updatedAt
						sourceLocale
						outdated
						outdatedLocales
						editedAt
						labelNames
						countryCode
						url
						htmlUrl
						name
						title
						body
						locale
						categoryConnection {
							id
							name
							keyName
						}
						sectionConnection {
							id
							name
						}
					}
				}
				`,
			},
			expectBody: map[string]interface{}{
				"data": map[string]interface{}{
					"oneArticle": map[string]interface{}{
						"id":              "115015959188",
						"authorId":        "24400224208",
						"commentsDisable": false,
						"draft":           false,
						"promoted":        false,
						"position":        0,
						"voteSum":         0,
						"voteCount":       0,
						"createdAt":       time.Date(2017, 12, 27, 3, 14, 16, 0, time.UTC),
						"updatedAt":       time.Date(2017, 12, 27, 3, 14, 48, 0, time.UTC),
						"sourceLocale":    "zh-tw",
						"outdated":        false,
						"outdatedLocales": []string{},
						"editedAt":        time.Date(2017, 12, 27, 3, 14, 16, 0, time.UTC),
						"labelNames":      []string{},
						"countryCode":     "tw",
						"url":             "https://honestbeehelp-tw.zendesk.com/api/v2/help_center/en-us/articles/115015959188-What-can-I-do-when-my-cart-is-locked-.json",
						"htmlUrl":         "https://help.honestbee.tw/hc/en-us/articles/115015959188-What-can-I-do-when-my-cart-is-locked-",
						"name":            "What can I do when my cart is locked?",
						"title":           "What can I do when my cart is locked?",
						"body":            "<p>When there is an error completing your checkout,your cart will be temporarily locked to prevent further changes to your order. To unlock your cart,click ‘Yes,unlock my cart,’ when prompted.</p>",
						"locale":          "en-us",
						"categoryConnection": map[string]interface{}{
							"id":      "115002432448",
							"name":    "My Account",
							"keyName": "myAccount",
						},
						"sectionConnection": map[string]interface{}{
							"id":   "115004118448",
							"name": "I need help with my account",
						},
					},
				},
			},
		},
		{
			description: "testing normal ZH_TW locale case",
			body: map[string]interface{}{
				"query": `
				{
					oneArticle(articleId: "115015959188", countryCode: TW, locale: ZH_TW) {
						id
						authorId
						commentsDisable
						draft
						promoted
						position
						voteSum
						voteCount
						createdAt
						updatedAt
						sourceLocale
						outdated
						outdatedLocales
						editedAt
						labelNames
						countryCode
						url
						htmlUrl
						name
						title
						body
						locale
						categoryConnection {
							id
							name
							keyName
						}
						sectionConnection {
							id
							name
						}
					}
				}
				`,
			},
			expectBody: map[string]interface{}{
				"data": map[string]interface{}{
					"oneArticle": map[string]interface{}{
						"id":              "115015959188",
						"authorId":        "24400224208",
						"commentsDisable": false,
						"draft":           false,
						"promoted":        false,
						"position":        0,
						"voteSum":         0,
						"voteCount":       0,
						"createdAt":       time.Date(2017, 12, 27, 3, 14, 16, 0, time.UTC),
						"updatedAt":       time.Date(2017, 12, 27, 3, 14, 48, 0, time.UTC),
						"sourceLocale":    "zh-tw",
						"outdated":        false,
						"outdatedLocales": []string{},
						"editedAt":        time.Date(2017, 12, 27, 3, 14, 16, 0, time.UTC),
						"labelNames":      []string{},
						"countryCode":     "tw",
						"url":             "https://honestbeehelp-tw.zendesk.com/api/v2/help_center/zh-tw/articles/115015959188-%E6%88%91%E7%9A%84%E8%B3%BC%E7%89%A9%E8%BB%8A%E9%8E%96%E4%BD%8F%E4%BA%86-%E8%A9%B2%E6%80%8E%E9%BA%BC%E8%BE%A6-.json",
						"htmlUrl":         "https://help.honestbee.tw/hc/zh-tw/articles/115015959188-%E6%88%91%E7%9A%84%E8%B3%BC%E7%89%A9%E8%BB%8A%E9%8E%96%E4%BD%8F%E4%BA%86-%E8%A9%B2%E6%80%8E%E9%BA%BC%E8%BE%A6-",
						"name":            "我的購物車鎖住了，該怎麼辦？",
						"title":           "我的購物車鎖住了，該怎麼辦？",
						"body":            "<p>當結帳出現錯誤時，您的購物車會暫時被鎖住，以避免您的訂單出現異動。要解鎖您的購物車，請在出現提示時，點選「是的，解鎖我的購物車」。</p>",
						"locale":          "zh-tw",
						"categoryConnection": map[string]interface{}{
							"id":      "115002432448",
							"name":    "我的帳號",
							"keyName": "myAccount",
						},
						"sectionConnection": map[string]interface{}{
							"id":   "115004118448",
							"name": "我需要帳號相關的協助",
						},
					},
				},
			},
		},
		{
			description: "testing article id not exist case",
			body: map[string]interface{}{
				"query": `
				{
					oneArticle(articleId: "334567833", countryCode: TW, locale: ZH_TW) {
						id
						authorId
						commentsDisable
						draft
						promoted
						position
						voteSum
						voteCount
						createdAt
						updatedAt
						sourceLocale
						outdated
						outdatedLocales
						editedAt
						labelNames
						countryCode
						url
						htmlUrl
						name
						title
						body
						locale
						categoryConnection {
							id
							name
							keyName
						}
						sectionConnection {
							id
							name
						}
					}
				}
				`,
			},
			expectBody: map[string]interface{}{
				"data": map[string]interface{}{
					"oneArticle": nil,
				},
				"errors": []interface{}{
					map[string]interface{}{
						"message": "Record Not Found",
						"path": []interface{}{
							"oneArticle",
						},
					},
				},
			},
		},
		{
			description: "testing not exist country code case",
			body: map[string]interface{}{
				"query": `
				{
					oneArticle(articleId: "115015959188", countryCode: not_exist_country_code) {
						id
						authorId
						commentsDisable
						draft
						promoted
						position
						voteSum
						voteCount
						createdAt
						updatedAt
						sourceLocale
						outdated
						outdatedLocales
						editedAt
						labelNames
						countryCode
						url
						htmlUrl
						name
						title
						body
						locale
						categoryConnection {
							id
							name
							keyName
						}
						sectionConnection {
							id
							name
						}
					}
				}
				`,
			},
			expectBody: map[string]interface{}{
				"errors": []interface{}{
					map[string]interface{}{
						"message": "Argument \"countryCode\" has invalid value not_exist_country_code.\nExpected type \"CountryCode\", found not_exist_country_code.",
						"locations": []interface{}{
							map[string]interface{}{
								"line":   3,
								"column": 57,
							},
						},
					},
				},
			},
		},
		{
			description: "testing not exist locale case",
			body: map[string]interface{}{
				"query": `
				{
					oneArticle(articleId: "115015959188", locale: not_exist_locale) {
						id
						authorId
						commentsDisable
						draft
						promoted
						position
						voteSum
						voteCount
						createdAt
						updatedAt
						sourceLocale
						outdated
						outdatedLocales
						editedAt
						labelNames
						countryCode
						url
						htmlUrl
						name
						title
						body
						locale
						categoryConnection {
						  id
						  name
						  keyName
						}
						sectionConnection {
						  id
						  name
						}
					}
				}
				`,
			},
			expectBody: map[string]interface{}{
				"errors": []interface{}{
					map[string]interface{}{
						"message": "Argument \"locale\" has invalid value not_exist_locale.\nExpected type \"Locale\", found not_exist_locale.",
						"locations": []interface{}{
							map[string]interface{}{
								"line":   3,
								"column": 52,
							},
						},
					},
				},
			},
		},
	}

	for _, tt := range testCases {
		t.Run(tt.description, func(t *testing.T) {
			// Send requests.
			b, err := json.Marshal(tt.body)
			if err != nil {
				t.Fatalf("[%s] json marshal failed:%v", tt.description, err)
			}
			resp, err := ts.Client().Post(ts.URL+"/graphql", "application/json", ioutil.NopCloser(bytes.NewReader(b)))
			if err != nil {
				t.Fatalf("[%s] http client get failed:%v", tt.description, err)
			}
			defer resp.Body.Close()

			// Compare HTTP status code.
			if http.StatusOK != resp.StatusCode {
				t.Errorf("[%s] http status expect:%v != actual:%v", tt.description, http.StatusOK, resp.StatusCode)
			}

			// Compare HTTP body.
			actual := make(map[string]interface{})
			if err = json.NewDecoder(resp.Body).Decode(&actual); err != nil {
				t.Fatalf("[%s] json decoding failed:%v", tt.description, err)
			}
			// Converts integer to the same type.
			expectData, err := json.Marshal(tt.expectBody)
			if err != nil {
				t.Fatalf("[%s] json marshal failed:%v", tt.description, err)
			}
			expect := make(map[string]interface{})
			if err = json.Unmarshal(expectData, &expect); err != nil {
				t.Fatalf("[%s] json unmarshal failed:%v", tt.description, err)
			}
			// Compares and prints difference.
			if diff := deep.Equal(expect, actual); diff != nil {
				t.Errorf("[%s] %v", tt.description, diff)
			}
		})
	}
}

func TestHandlersGraphQLQuerySearchTitleArticles(t *testing.T) {
}

func TestHandlersGraphQLQuerySearchBodyArticles(t *testing.T) {
}
