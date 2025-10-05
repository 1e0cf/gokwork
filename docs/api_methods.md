# API Methods Reference

## createAnswer

**Path:** `POST /createAnswer`

**Summary:** Создать ответ на отзыв

---

## editAnswer

**Path:** `POST /editAnswer`

**Summary:** Редактировать отзыв

---

## rechargeBalance

**Path:** `POST /rechargeBalance`

**Summary:** Получить ссылку для пополнения баланса в профиле и для создания заказа.

**Description:** Пример доступных типов пополнения:</br>
"paypal": "PayPal",</br>
"paymore_card": "Банковская карта",</br>
"bank": "Безнал для юрлиц и ИП",</br>
"paymore_webmoney": "WebMoney",</br>
"webmoney": "WebMoney" - напрямую,</br>
"paymore_yandex": "ЮMoney",</br>
"paymore_qiwi": "QIWI Кошелек"</br>

---

## getPaymentMethods

**Path:** `POST /getPaymentMethods`

**Summary:** Получение способов оплаты и информации о сервисном сборе

---

## getBillRefillUrl

**Path:** `POST /getBillRefillUrl`

**Summary:** Получить ссылку для Yescrow платежа

---

## markKworksBlackFriday

**Path:** `POST /markKworksBlackFriday`

**Summary:** Отмечает кворк участвующим в акции Черная пятница

**Description:** Отмечает кворк участвующим в акции Черная пятница

---

## catalogFilters

**Path:** `POST /catalogFilters`

**Summary:** Список фильтров из текущего раздела каталога

---

## categories

**Path:** `POST /categories`

**Summary:** Список категорий

---

## translationLanguages

**Path:** `POST /translationLanguages`

**Summary:** Получить массив всех доступных в системе языков для переводов с падежами

---

## categoryAttributes

**Path:** `POST /categoryAttributes`

**Summary:** Получить дерево всех атрибутов для заданной категории

---

## positiveReviewsCount

**Path:** `POST /positiveReviewsCount`

**Summary:** Получить массив количества отзывов по атрибуту

---

## category

**Path:** `POST /category`

**Summary:** Получить данные о заданной категории

---

## cities

**Path:** `POST /cities`

**Summary:** Получение списка городов страны

---

## registerCloudToken

**Path:** `POST /registerCloudToken`

**Summary:** Регистрация токена Firebase Cloud Messaging

---

## fcmTokenRequestFailed

**Path:** `POST /fcmTokenRequestFailed`

**Summary:** Регистрация токена Firebase Cloud Messaging

---

## fcmNotificationsReceived

**Path:** `POST /fcmNotificationsReceived`

**Summary:** Пометка сообщений FCM полученными в МП

---

## fcmNotificationsRead

**Path:** `POST /fcmNotificationsRead`

**Summary:** Пометка сообщений FCM прочитанными в МП

---

## getComplainCategories

**Path:** `POST /getComplainCategories`

**Summary:** Получение списка жалоб

---

## createKworkComplain

**Path:** `POST /createKworkComplain`

**Summary:** Создание жалобы на кворк

---

## countries

**Path:** `POST /countries`

**Summary:** Получение списка стран

---

## dialogs

**Path:** `POST /dialogs`

**Summary:** Список диалогов

---

## blockedDialogList

**Path:** `POST /blockedDialogList`

**Summary:** Список заблокированных для диалога юзеров

---

## blockDialog

**Path:** `POST /blockDialog`

**Summary:** Заблокировать диалог

---

## unblockDialog

**Path:** `POST /unblockDialog`

**Summary:** Разблокировать диалог

---

## archiveDialog

**Path:** `POST /archiveDialog`

**Summary:** Отправить диалог в архив

---

## unarchiveDialog

**Path:** `POST /unarchiveDialog`

**Summary:** Вернуть диалог из архива

---

## setDialogStarred

**Path:** `POST /setDialogStarred`

**Summary:** Пометить диалог избранным

---

## unreadDialog

**Path:** `POST /unreadDialog`

**Summary:** Пометить диалог с заданным пользователем непрочитанным

---

## searchDialogs

**Path:** `POST /searchDialogs`

**Summary:** Поиск диалогов

---

## hideDialog

**Path:** `POST /hideDialog`

**Summary:** Скрыть/удалить диалог

---

## getDialog

**Path:** `POST /getDialog`

**Summary:** Получить диалог по идентификатору собеседника

---

## isDialogAllow

**Path:** `POST /isDialogAllow`

**Summary:** Разрешен ли диалог с пользователем

---

## workerConfirmsExtraRemovalRequest

**Path:** `POST /workerConfirmsExtraRemovalRequest`

**Summary:** Продавец подтверждает запрос на удаление опции

---

## payerDeclinesExtraRemovalRequest

**Path:** `POST /payerDeclinesExtraRemovalRequest`

**Summary:** Покупатель отклоняет запрос на удаление опции из заказа

---

## workerDeclinesExtraRemovalRequest

**Path:** `POST /workerDeclinesExtraRemovalRequest`

**Summary:** Продавец отклоняет запрос на удаление опции из заказа

---

## workerExtraDelete

**Path:** `POST /workerExtraDelete`

**Summary:** Удалить опцию из заказа, для продавца

---

## payerExtraDelete

**Path:** `POST /payerExtraDelete`

**Summary:** Удалить опцию из заказа, для покупателя

---

## payerBuyExtras

**Path:** `POST /payerBuyExtras`

**Summary:** Добавление опции в заказ покупателем

---

## acceptExtras

**Path:** `POST /acceptExtras`

**Summary:** Принятие предложенных опции в треке покупателем

---

## getAvailableFeatures

**Path:** `POST /getAvailableFeatures`

**Summary:** Информация о доступных функциях

---

## setVoiceMessageSpeed

**Path:** `POST /setVoiceMessageSpeed`

**Summary:** Изменение скорости воспроизведения голосовых сообщений

---

## hideVoiceMessageSettingsPopup

**Path:** `POST /hideVoiceMessageSettingsPopup`

**Summary:** Отправка на бэк факта показа попапа

---

## setVoiceMessageReceiving

**Path:** `POST /setVoiceMessageReceiving`

**Summary:** Разрешить/запретить принимать голосовые сообщения

---

## fileUpload

**Path:** `POST /fileUpload`

**Summary:** Загрузка файла из FILES["upload_files"]

---

## voiceUpload

**Path:** `POST /voiceUpload`

**Summary:** Загрузка файла из FILES["upload_files"]

---

## fileDelete

**Path:** `POST /fileDelete`

**Summary:** Удаление загруженного файла

---

## uploadLog

**Path:** `POST /uploadLog`

**Summary:** Загрузка лога мобильного приложения

---

## uploadedFile

**Path:** `POST /uploadedFile`

**Summary:** Получение загруженного файла

---

## miniature

**Path:** `POST /miniature`

**Summary:** Получить миниатюру к файлу

---

## getFishingTutorialQuestions

**Path:** `POST /getFishingTutorialQuestions`

**Summary:** Список вопросов для опроса о мошенниках (код 116)

---

## setFishingTutorialStatus

**Path:** `POST /setFishingTutorialStatus`

**Summary:** Установка статуса о прохождении опроса о мошенниках

---

## inboxes

**Path:** `POST /inboxes`

**Summary:** Список сообщений в диалоге

---

## getInboxTracks

**Path:** `POST /getInboxTracks`

**Summary:** Список сообщений (c треками) в диалоге

---

## markInboxTracksAsRead

**Path:** `POST /markInboxTracksAsRead`

**Summary:** Пометить переписку пользователя прочитанной

---

## searchInboxes

**Path:** `POST /searchInboxes`

**Summary:** Поиск сообщенений в диалогах пользователей

---

## searchMessages

**Path:** `POST /searchMessages`

**Summary:** Поиск сообщенений в диалогах пользователей с указанием найденного сниппета

---

## inboxDelete

**Path:** `POST /inboxDelete`

**Summary:** Удаление сообщения

---

## inboxEdit

**Path:** `POST /inboxEdit`

**Summary:** Редактирование сообщения

---

## inboxCreate

**Path:** `POST /inboxCreate`

**Summary:** Создание нового сообщения

---

## inboxComplainMessage

**Path:** `POST /inboxComplainMessage`

**Summary:** Пожаловаться на сообщение пользователя

---

## inboxForward

**Path:** `POST /inboxForward`

**Summary:** Пересылка сообщения

---

## inboxRead

**Path:** `POST /inboxRead`

**Summary:** Пометить прочитанным сообщения или диалог

---

## inboxMessage

**Path:** `POST /inboxMessage`

**Summary:** Получить сообщение Inbox по id

---

## inboxTrackMessage

**Path:** `POST /inboxTrackMessage`

**Summary:** Получить сообщение Inbox/Track по conversationId

---

## typing

**Path:** `POST /typing`

**Summary:** Отправить флаг "Юзер печатает"

---

## sendUserStatus

**Path:** `POST /sendUserStatus`

**Summary:** Отправить флаг "Статус пользователя"

---

## markVoiceMessageHeard

**Path:** `POST /markVoiceMessageHeard`

**Summary:** Отправить флаг "Пользователь прослушал голосовое сообщение"

---

## allowInboxRequest

**Path:** `POST /allowInboxRequest`

**Summary:** Разрешить/не разрешить переписку с пользователем

---

## getVoiceMessageTranscription

**Path:** `POST /getVoiceMessageTranscription`

**Summary:** Получить транскрипцию голосового сообщения

---

## getVoiceMessageConvertStatus

**Path:** `POST /getVoiceMessageConvertStatus`

**Summary:** Получить статус конвертации голосового сообщения

---

## updateChatDraftMessage

**Path:** `POST /updateChatDraftMessage`

**Summary:** Обновление черновика

---

## inboxCustomRequestDecline

**Path:** `POST /inboxCustomRequestDecline`

**Summary:** Отмена индивидуального запроса

---

## inboxPayerDecline

**Path:** `POST /inboxPayerDecline`

**Summary:** Отмена предложенного кворка в личке покупателем

---

## inboxWorkerDecline

**Path:** `POST /inboxWorkerDecline`

**Summary:** Отмена предложенного кворка в личке продавцом

---

## kworks

**Path:** `POST /kworks`

**Summary:** Список кворков для категории

---

## userKworks

**Path:** `POST /userKworks`

**Summary:** Список кворков пользователя

---

## getHiddenKworks

**Path:** `POST /getHiddenKworks`

**Summary:** Получить список скрытых пользователем кворков

---

## kworksStatusList

**Path:** `POST /kworksStatusList`

**Summary:** Получение статусов (вкладок) кворков, и первую страницу самих кворков для каждого статуса

---

## markKworkAsHidden

**Path:** `POST /markKworkAsHidden`

**Summary:** Скрытие\отображение кворка

---

## markKworkAsFavorite

**Path:** `POST /markKworkAsFavorite`

**Summary:** Добавление кворка в избранные или удаление из избранных

---

## pauseKwork

**Path:** `POST /pauseKwork`

**Summary:** Останавливает кворк

---

## deleteKwork

**Path:** `POST /deleteKwork`

**Summary:** Удаление своего кворка

---

## startKwork

**Path:** `POST /startKwork`

**Summary:** Активирует(запускает) кворк

---

## viewedCatalogKworks

**Path:** `POST /viewedCatalogKworks`

**Summary:** Список просмотренных кворков

---

## favoriteKworks

**Path:** `POST /favoriteKworks`

**Summary:** Список избранных кворков

---

## getKworkDetails

**Path:** `POST /getKworkDetails`

**Summary:** Получение данных о кворке

---

## getSubscribersStatistics

**Path:** `POST /getSubscribersStatistics`

**Summary:** Данные о динамике подписчиков на канале

---

## getKworkDetailsExtra

**Path:** `POST /getKworkDetailsExtra`

**Summary:** Получение данных о кворке

---

## getKworkAnswers

**Path:** `POST /getKworkAnswers`

**Summary:** Получение FAQ по кворку

---

## getKworkLinksTable

**Path:** `POST /getKworkLinksTable`

**Summary:** Получить данные по ссылкам кворка

---

## getKworkLinksTablev2

**Path:** `POST /getKworkLinksTablev2`

**Summary:** Получить данные по ссылкам кворка

---

## getKworkPortfolios

**Path:** `POST /getKworkPortfolios`

**Summary:** Получить работы портфолио для кворка

---

## getKworkReviews

**Path:** `POST /getKworkReviews`

**Summary:** Получить отзывы для кворка

---

## catalogMain

**Path:** `POST /catalogMain`

**Summary:** Получение основной информации по главной странице каталога

**Description:** Получение основной информации по главной странице каталога

---

## catalogRubrics

**Path:** `POST /catalogRubrics`

**Summary:** Рубрики меню

---

## catalogCategories

**Path:** `POST /catalogCategories`

**Summary:** Получение списка подкатегорий

---

## getPayerCompanyModalUrl

**Path:** `GET /getPayerCompanyModalUrl`

**Summary:** Получение ссылки на страницу покупателя с открытой модалкой регистрации Компании

**Description:** Получение ссылки на страницу покупателя с открытой модалкой регистрации Компании

---

## catalogMainv2

**Path:** `POST /catalogMainv2`

**Summary:** Получение основной информации по главной странице нового каталога

**Description:** Получение основной информации по главной странице нового каталога

---

## notifications

**Path:** `POST /notifications`

**Summary:** Список уведомлений пользователя

---

## notificationsReceived

**Path:** `POST /notificationsReceived`

**Summary:** Пометка сквозных Push-событий прочтенными

---

## notificationsFetch

**Path:** `POST /notificationsFetch`

**Summary:** Получение непрочитанных Push-событий из очереди

---

## pushInAppNotificationLog

**Path:** `POST /pushInAppNotificationLog`

**Summary:** Запись лога показа in-app уведомления

---

## offer

**Path:** `POST /offer`

**Summary:** Получает данные предложения к запросу на услугу на бирже

---

## offers

**Path:** `POST /offers`

**Summary:** Предложения пользователя на запросы услуг на бирже

---

## deleteOffer

**Path:** `POST /deleteOffer`

**Summary:** Удаляет предложение к запросу на услугу на бирже

---

## order

**Path:** `POST /order`

**Summary:** Информация о заказе

---

## getOrderFiles

**Path:** `POST /getOrderFiles`

**Summary:** Получить список файлов заказа

---

## getOrderHeader

**Path:** `POST /getOrderHeader`

**Summary:** Обновление блоков кеша шапки заказа

---

## ordersBetween

**Path:** `POST /ordersBetween`

**Summary:** Список активных заказов между текущим пользователем и указанным

---

## getUsersLastOrderInfo

**Path:** `POST /getUsersLastOrderInfo`

**Summary:** Возвращает информацию по последнему выполненному заказу между пользователем и собеседником, в котором было хотя бы одно сообщение

---

## orderKwork

**Path:** `POST /orderKwork`

**Summary:** Создать заказ по кворку.

**Description:** Виды заказов (ожидаемые параметры):
1. Беспакетный с ценой за кворк:
$kworkCount(required) - количество кворков
$kworkId(required) - идентификатор кворка
$extras - опциональный набор покупаемых опций
$orderId - обязательный только в случае когда создется новый инд. заказ по содержимому выполненного инд. заказа.

2. Беспакетный с ценой за фиксированный объем
$kworkId(required) - идентификатор кворка
$volumeTypeId - опциональный идентификатор типа объема, для случая если кворк имеет доп.объем
$extras - опциональный набор покупаемых опций
$volume(require) - купленный объем
$orderId - обязательный только в случае когда создется новый инд. заказ по содержимому выполненного инд. заказа.

3. Однопакетный с ценой за кворк
$kworkId(required) - идентификатор кворка
$kworkCount(required) - количество кворков
$extras - опциональный набор покупаемых опций
$orderId - обязательный только в случае когда создется новый инд. заказ по содержимому выполненного инд. заказа.

4. Однопакетный с ценой за фиксированный объем
$kworkId(required) - идентификатор кворка
$volumeTypeId - опциональный идентификатор типа объема, для случая если кворк имеет доп.объем
$extras - опциональный набор покупаемых опций
$volume(require) - купленный объем
$orderId - обязательный только в случае когда создется новый инд. заказ по содержимому выполненного инд. заказа.
$packageId(require) - идентификатор пакета

5. Многопакетный с ценой за кворк
$kworkId(required) - идентификатор кворка
$orderId - обязательный только в случае когда создется новый инд. заказ по содержимому выполненного инд. заказа.
$packageId(require) - идентификатор пакета

6. Многопакетный с ценой за фиксированный объем
$kworkId(required) - идентификатор кворка
$volumeTypeId - опциональный идентификатор типа объема, для случая если кворк имеет доп.объем
$volume(require) - купленный объем
$orderId - обязательный только в случае когда создется новый инд. заказ по содержимому выполненного инд. заказа.
$packageId(require) - идентификатор пакета

---

## repeatOrder

**Path:** `POST /repeatOrder`

**Summary:** Создать заказ заново

---

## getTracks

**Path:** `POST /getTracks`

**Summary:** Возвращает информацию о треках заказа

---

## getOrderDetails

**Path:** `POST /getOrderDetails`

**Summary:** Получение детальной информации о заказе

---

## sendOrderForApproval

**Path:** `POST /sendOrderForApproval`

**Summary:** Отправка заказа на проверку

---

## approveOrder

**Path:** `POST /approveOrder`

**Summary:** Принятие заказа

---

## approveOrderStage

**Path:** `POST /approveOrderStage`

**Summary:** Принятие этапов заказа

---

## getExtrasAvailableForOrder

**Path:** `POST /getExtrasAvailableForOrder`

**Summary:** Получить опции, доступные для добавления в заказ

---

## getOrderedExtras

**Path:** `POST /getOrderedExtras`

**Summary:** Получение дополнительных опций, которые уже добавлены в заказ

---

## getCustomOptionsPresets

**Path:** `POST /getCustomOptionsPresets`

**Summary:** Получить спиок цен для кастомных опций и добавляемого срока

---

## getOrderCancellationReasons

**Path:** `POST /getOrderCancellationReasons`

**Summary:** Получить список причин отмены заказа

---

## cancelOrderAwaitingPayment

**Path:** `POST /cancelOrderAwaitingPayment`

**Summary:** Удалить неоплаченный заказ

---

## sendOrderForRevision

**Path:** `POST /sendOrderForRevision`

**Summary:** Отправить заказ на доработку

---

## sendBonus

**Path:** `POST /sendBonus`

**Summary:** Отправить бонус продавцу

---

## getArbitrationReasons

**Path:** `POST /getArbitrationReasons`

**Summary:** Получить список причин перевода в арбитраж

---

## sendOrderForArbitration

**Path:** `POST /sendOrderForArbitration`

**Summary:** Отправить заказ на арбитраж

---

## offerOrderOptions

**Path:** `POST /offerOrderOptions`

**Summary:** Добавить предложение опций к заказу

---

## sendOrderRequirements

**Path:** `POST /sendOrderRequirements`

**Summary:** Отправка информации по заказу продавцу

---

## sendReport

**Path:** `POST /sendReport`

**Summary:** Отправить отчет по заказу (не этапному)

---

## saveOrderNote

**Path:** `POST /saveOrderNote`

**Summary:** Создать/редактировать заметку о заказе

---

## deleteOrderNote

**Path:** `POST /deleteOrderNote`

**Summary:** Удалить заметку о заказе

---

## payOrderAwaitingPayment

**Path:** `POST /payOrderAwaitingPayment`

**Summary:** Оплата заказа в статусе "Ожидает оплаты"

---

## allowOrderPortfolioUpload

**Path:** `POST /allowOrderPortfolioUpload`

**Summary:** Разрешить продавцу публикацию работы в портфолио

---

## getOrderProvidedData

**Path:** `POST /getOrderProvidedData`

**Summary:** Предоставленные данные по заказу

---

## setOrderRating

**Path:** `POST /setOrderRating`

**Summary:** Оценить продавца

---

## sendOrderReceiptLinkForVerification

**Path:** `POST /sendOrderReceiptLinkForVerification`

**Summary:** Отправка ссылки на чек для проверки

---

## rateArbitration

**Path:** `POST /rateArbitration`

**Summary:** Выставление оценки за арбитраж

---

## createStage

**Path:** `POST /createStage`

**Summary:** Добавление этапа в заказ

---

## addStage

**Path:** `POST /addStage`

**Summary:** Создание и резервирование этапа в заказе

---

## orderStage

**Path:** `POST /orderStage`

**Summary:** Зарезервировать этапы

---

## updateStageProgress

**Path:** `POST /updateStageProgress`

**Summary:** Обновить прогресс по задаче

---

## editStage

**Path:** `POST /editStage`

**Summary:** Редактирование этапов в заказе

---

## deleteStage

**Path:** `POST /deleteStage`

**Summary:** Удаление этапа из заказа

---

## suggestStages

**Path:** `POST /suggestStages`

**Summary:** Встречное предложение этапов

---

## rejectStageSuggestion

**Path:** `POST /rejectStageSuggestion`

**Summary:** Отмена встречного предложения этапов

---

## acceptStageSuggestion

**Path:** `POST /acceptStageSuggestion`

**Summary:** Принятие встречного предложения этапов

---

## payerUpgradePackage

**Path:** `POST /payerUpgradePackage`

**Summary:** Апгрейд пакета покупателем

---

## payerOrders

**Path:** `POST /payerOrders`

**Summary:** Список заказов покупателя

---

## cancelOrderByPayer

**Path:** `POST /cancelOrderByPayer`

**Summary:** Покупатель отменяет заказ

---

## confirmCancelOrderRequestByPayer

**Path:** `POST /confirmCancelOrderRequestByPayer`

**Summary:** Покупатель соглашается на обоюдную отмену заказа

---

## rejectCancelOrderRequestByPayer

**Path:** `POST /rejectCancelOrderRequestByPayer`

**Summary:** Покупатель не согласился на обоюдную отмену заказа

---

## deleteCancelOrderRequestByPayer

**Path:** `POST /deleteCancelOrderRequestByPayer`

**Summary:** Покупатель удалил свой запрос на обоюдную отмену заказа

---

## payerDeclineExtras

**Path:** `POST /payerDeclineExtras`

**Summary:** Покупатель отклоняет предложенные опции

---

## editPortfolio

**Path:** `POST /editPortfolio`

**Summary:** Редактировать портфолио

---

## createPortfolio

**Path:** `POST /createPortfolio`

**Summary:** Добавить портфолио

---

## portfolioList

**Path:** `POST /portfolioList`

**Summary:** Получить портфолио пользователя

---

## portfolioCategoriesList

**Path:** `POST /portfolioCategoriesList`

**Summary:** Получение категорий работ, и первой страницы самих работ для каждой категории

---

## deletePortfolio

**Path:** `POST /deletePortfolio`

**Summary:** Удаление портфолио

---

## uploadPortfolioFile

**Path:** `POST /uploadPortfolioFile`

**Summary:** Загрузка файла из FILES["file"]

---

## createReview

**Path:** `POST /createReview`

**Summary:** Создать отзыв

---

## editReview

**Path:** `POST /editReview`

**Summary:** Редактировать отзыв

---

## deleteReview

**Path:** `POST /deleteReview`

**Summary:** Удаление отзыва

---

## search

**Path:** `POST /search`

**Summary:** Поиск кворков

---

## userSearch

**Path:** `POST /userSearch`

**Summary:** Поиск пользователей

---

## searchKworksCatalogQuery

**Path:** `POST /searchKworksCatalogQuery`

**Summary:** Получение ключевых слов по частичной фразе

---

## deleteAccount

**Path:** `POST /deleteAccount`

**Summary:** Удаление своего аккаунта

---

## verifySmsCodeForAccountDeleting

**Path:** `POST /verifySmsCodeForAccountDeleting`

**Summary:** Проверка кода удаления аккаунта

---

## sendSelfEmployedSurveyResult

**Path:** `POST /sendSelfEmployedSurveyResult`

**Summary:** Отправить ответ на опрос

---

## hideSelfEmployedNotification

**Path:** `POST /hideSelfEmployedNotification`

**Summary:** Скрыть уведомление об успешной регистрации СЗ/ИП

---

## terms

**Path:** `POST /terms`

**Summary:** Вывод договора-оферты

---

## privacy

**Path:** `POST /privacy`

**Summary:** Вывод политики конфиденциальности

---

## termsOfService

**Path:** `POST /termsOfService`

**Summary:** Вывод правил сайта

---

## resolution

**Path:** `POST /resolution`

**Summary:** Вывод политики разрешения споров

---

## timezones

**Path:** `POST /timezones`

**Summary:** Получение списка временных зон

---

## trackMessage

**Path:** `POST /trackMessage`

**Summary:** Получить сообщение трека

---

## trackDelete

**Path:** `POST /trackDelete`

**Summary:** Удаление трека

---

## trackEdit

**Path:** `POST /trackEdit`

**Summary:** Редактирование трека

---

## trackRead

**Path:** `POST /trackRead`

**Summary:** Пометить указанные треки прочитанными

---

## searchTracks

**Path:** `POST /searchTracks`

**Summary:** Поиск текстовых треков в заказе

---

## searchOrderTracks

**Path:** `POST /searchOrderTracks`

**Summary:** Поиск текстовых треков в заказе

---

## updateOrderDraftMessage

**Path:** `POST /updateOrderDraftMessage`

**Summary:** Обновление черновика

---

## signIn

**Path:** `POST /signIn`

**Summary:** Аутентификация пользователя с выдачей токена

**Description:** Если пришел код ошибки 118 то необходимо показать пользователю webview со страницей http://kwork.ru/captcha_only.
			<br/>При успешном прохождении капчи к url страницы будет добавлен хеш вида #response=xxx,
			где xxx - данные которые необходимо послать в параметре g-recaptcha-response

---

## logout

**Path:** `POST /logout`

**Summary:** Выход пользователя: удаление указанного пуш токена, удаление токена авторизации, закрытие текущей сессии

---

## signUp

**Path:** `POST /signUp`

**Summary:** Регистрация пользователя с выдачей токена

**Description:** Если пришел код ошибки 118 то необходимо показать пользователю webview со страницей http://kwork.ru/captcha_only.
			<br/>При успешном прохождении капчи к url страницы будет добавлен хеш вида #response=xxx,
			где xxx - данные которые необходимо послать в параметре g-recaptcha-response

---

## socialSignIn

**Path:** `POST /socialSignIn`

**Summary:** Аутентификация пользователя через социальные сети по коду провайдера

**Description:** Успешная авторизация только в случае если пользователь дал доступ к своему email и у него есть email,
			если в данных пользователя нет email то будет ошибка с кодом 159, которая указывает что необходимо воспользоваться методом socialSignUp

---

## socialSignInByToken

**Path:** `POST /socialSignInByToken`

**Summary:** Аутентификация пользователя через токен социальной сети (при нативной авторизации моб. приложений)
Аутентификация + регистрация
(Устаревший метод, вскоре должен быть удален)

**Description:** Успешная авторизация только в случае если пользователь дал доступ к своему email и у него есть email,
			если в данных пользователя нет email то будет ошибка с кодом 159, которая указывает что необходимо воспользоваться методом socialSignUp

---

## socialSignInByTokenv2

**Path:** `POST /socialSignInByTokenv2`

**Summary:** Аутентификация пользователя через токен социальной сети (при нативной авторизации моб. приложений)

**Description:** Если пользователь не найден, то возвращается ошибка с кодом 159, которая указывает что необходимо воспользоваться методом socialSignUp.

---

## socialSignUp

**Path:** `POST /socialSignUp`

**Summary:** Регистрация через социальные сети с указанием email (привязка социального аккаунта если такой email уже есть)

**Description:** Использовать его необходимо только после метода socialSignIn, если тот вернул ошибку с кодом 159.
			Успешная регистрация в случае если пользователя с таким email нет среди пользователей kwork, если пользователь существует то будет ошибка с кодом 160,
			а также будет выслано письмо на указанный email для подтверждения привязки социального аккаунта к аккаунту kwork

---

## socialSignUpByToken

**Path:** `POST /socialSignUpByToken`

**Summary:** Регистрация через социальные сети по токену с указанием email (привязка социального аккаунта если такой email уже есть)

**Description:** Использовать его необходимо только после метода socialSignInByToken, если тот вернул ошибку с кодом 159.
			Успешная регистрация в случае если пользователя с таким email нет среди пользователей kwork, если пользователь существует то будет ошибка с кодом 160,
			а также будет выслано письмо на указанный email для подтверждения привязки социального аккаунта к аккаунту kwork

---

## appleSignIn

**Path:** `POST /appleSignIn`

**Summary:** Аутентификация пользователя через Apple

**Description:** Успешная авторизация только в случае если пользователь дал доступ к своему email и нам его отдали,
	 * если в данных пользователя нет email то будет ошибка с кодом 159, которая указывает что необходимо воспользоваться другим методом авторизации

---

## actor

**Path:** `POST /actor`

**Summary:** Данные авторизованного пользователя

---

## getBadgesInfo

**Path:** `POST /getBadgesInfo`

**Summary:** Плучение информации для бейджей о количестве важных уведомлений

---

## kworksCategoriesList

**Path:** `POST /kworksCategoriesList`

**Summary:** Получение вкладок пользователя с категориями кворков

---

## user

**Path:** `POST /user`

**Summary:** Данные пользователя по идентификатору

**Description:** Публичные данные пользователя

---

## getActorInfo

**Path:** `POST /getActorInfo`

**Summary:** Информация о текущем залогиненном пользователе

---

## getSecurityUserData

**Path:** `POST /getSecurityUserData`

**Summary:** Получения данных для экрана "Безопасность"

---

## getUserInfo

**Path:** `POST /getUserInfo`

**Summary:** Получение информации о пользователе

---

## setTakingOrders

**Path:** `POST /setTakingOrders`

**Summary:** Сохранение настройки пользователя по доступности его кворков для заказа

---

## userByUsername

**Path:** `POST /userByUsername`

**Summary:** Получение данных пользователя по username

---

## changePassword

**Path:** `POST /changePassword`

**Summary:** Запрос изменения пароля пользователя, с отправкой посылка кода подтверждения на почту

---

## updateSettings

**Path:** `POST /updateSettings`

**Summary:** Редактирование настроек пользователя

**Description:** Редактирование настроек пользователя
Помимо указанных параметров ожидаются загруженные файлы
$_FILES["avatar-photo"], остаётся в этом методе для поддержки старых сборок

---

## updateAvatar

**Path:** `POST /updateAvatar`

**Summary:** Смена аватара пользователя

**Description:** Смена аватара пользователя

---

## uploadCover

**Path:** `POST /uploadCover`

**Summary:** Загрузка обложки пользователя

**Description:** Загрузка обложки пользователя

---

## deleteCover

**Path:** `POST /deleteCover`

**Summary:** Удаление обложки пользователя

**Description:** Удаление обложки пользователя

---

## resetPassword

**Path:** `POST /resetPassword`

**Summary:** Отправка письма для сброса пароля

---

## getCookie

**Path:** `POST /getCookie`

**Summary:** Получение куки для авторизованных по токену

---

## getChannel

**Path:** `POST /getChannel`

**Summary:** Получить идентификатор socket-канала текущего пользователя

---

## offline

**Path:** `POST /offline`

**Summary:** Установить статус offline

---

## allowMobilePush

**Path:** `POST /allowMobilePush`

**Summary:** Устанавливает/снимает флаг разрешения отправки пуша в мобильное приложение

---

## setAvailableAtWeekends

**Path:** `POST /setAvailableAtWeekends`

**Summary:** Изменение доступности кворков на выходных

**Description:** Изменение доступности кворков на выходных

---

## setUserType

**Path:** `POST /setUserType`

**Summary:** Установка типа пользователя (покупатель/продавец)

**Description:** Установка типа пользователя (покупатель/продавец)

---

## allowPushNotificationsSound

**Path:** `POST /allowPushNotificationsSound`

**Summary:** Устанавливать / снимать флаг разрешения воспроизведения звука при отображении пуша в мобильном приложении.

**Description:** Установка настройки звука в мобильном приложении

---

## reportAppVersion

**Path:** `POST /reportAppVersion`

**Summary:** Обновление версии мобильного приложения пользователя

**Description:** Обновление версии мобильного приложения пользователя

---

## emailVerificationLetter

**Path:** `POST /emailVerificationLetter`

**Summary:** Запрос отправки письма подтверждения email

**Description:** Запрос отправки письма подтверждения email

---

## saveUserNote

**Path:** `POST /saveUserNote`

**Summary:** Создать/редактировать заметку о пользователе

---

## deleteUserNote

**Path:** `POST /deleteUserNote`

**Summary:** Удалить заметку о пользователе

---

## getCaptchaStatus

**Path:** `POST /getCaptchaStatus`

**Summary:** Требуется ли показать капчу при запросе сброса пароля

---

## changePayerSubRole

**Path:** `POST /changePayerSubRole`

**Summary:** Сменить дочернюю роль покупателя

---

## changeUsername

**Path:** `POST /changeUsername`

**Summary:** Запрос изменения логина пользователя

---

## checkLogin

**Path:** `POST /checkLogin`

**Summary:** Запрос проверки занятости логина

---

## favoriteCategories

**Path:** `POST /favoriteCategories`

**Summary:** Получение списка любимых рубрик пользователя.

---

## addFavoriteCategories

**Path:** `POST /addFavoriteCategories`

**Summary:** Изменение списка любимых категорий пользователя

---

## delFavoriteCategories

**Path:** `POST /delFavoriteCategories`

**Summary:** Удаление любимых категорий пользователя

---

## setFavorite

**Path:** `POST /setFavorite`

**Summary:** Изменение списка любимых категорий пользователя, объединяет функционал add и delete

---

## addPhoneNumber

**Path:** `POST /addPhoneNumber`

**Summary:** Запрос кода для добавления номера телефона

**Description:** Запрос кода для добавления номера телефона

---

## verifyPhoneActivationCode

**Path:** `POST /verifyPhoneActivationCode`

**Summary:** Проверка кода активации номера телефона

**Description:** Проверка кода активации номера телефона

---

## requestPhoneChanging

**Path:** `POST /requestPhoneChanging`

**Summary:** Запрос кода для смены номера телефона

**Description:** Запрос кода для смены номера телефона

---

## addNewPhoneNumber

**Path:** `POST /addNewPhoneNumber`

**Summary:** Проверка кода активации для добавления нового номера телефона на замену старому

**Description:** Проверка валидационного хеша для добавления нового номера телефона на замену старому

---

## sendWhatsAppCode

**Path:** `POST /sendWhatsAppCode`

**Summary:** Отправить код верификации через WhatsApp

**Description:** Отправка кода по WhatsApp

---

## userReviews

**Path:** `POST /userReviews`

**Summary:** Список отзывов для пользователя

---

## getCurrentVersions

**Path:** `POST /getCurrentVersions`

**Summary:** Возвращает список текущих версий мобильных приложений

---

## myWants

**Path:** `POST /myWants`

**Summary:** Возвращает список запросов на услугу

---

## want

**Path:** `POST /want`

**Summary:** Возвращает данные по запросу на услугу

---

## restartWant

**Path:** `POST /restartWant`

**Summary:** Перезапускает запрос на услугу на бирже

---

## stopWant

**Path:** `POST /stopWant`

**Summary:** Останавливает запрос на услугу на бирже

---

## deleteWant

**Path:** `POST /deleteWant`

**Summary:** Удаляет запрос на услугу на бирже

---

## wantsStatusList

**Path:** `POST /wantsStatusList`

**Summary:** Список запросов на услугу на бирже, сгруппированных по альтернативному статусу

---

## exchangeInfo

**Path:** `POST /exchangeInfo`

**Summary:** Ключевая информации по бирже

---

## projects

**Path:** `POST /projects`

**Summary:** Возвращает список проектов по заданным фильтрам

---

## getWantsCount

**Path:** `POST /getWantsCount`

**Summary:** Возвращает количество проектов по заданным фильтрам

---

## applyFilters

**Path:** `POST /applyFilters`

**Summary:** Установить выбранные фильтры продавца на бирже

---

## clearFilters

**Path:** `POST /clearFilters`

**Summary:** Сбросить выбранных фильтров продавца на бирже

---

## getWebAuthToken

**Path:** `POST /getWebAuthToken`

**Summary:** Получить токен входа в веб версию

**Description:** Возвращает токен со сроком жизни, а также готовую ссылку с токеном для логина в веб версии.
			<br/>Юзер может создать сколько угодно токенов, каждый будет одноразовым. Время жизни одна минута.
			<br/>Для входа в веб версию существует роут https://kwork.ru/login-by-token, который ожидает GET параметр web_auth_token с полученным токеном
			(полный URL возвращается в response.url).
			<br/>Можно передать опциональный параметр url_to_redirect, если требуется, чтобы после перехода в веб версию открывалась определенная страница.

---

## workerOrders

**Path:** `POST /workerOrders`

**Summary:** Список заказов пользователя, которые он должен выполнить

---

## cancelOrderByWorker

**Path:** `POST /cancelOrderByWorker`

**Summary:** Продавец отменяет заказ

---

## confirmCancelOrderRequestByWorker

**Path:** `POST /confirmCancelOrderRequestByWorker`

**Summary:** Продавец соглашается на обоюдную отмену заказа

---

## rejectCancelOrderRequestByWorker

**Path:** `POST /rejectCancelOrderRequestByWorker`

**Summary:** Продавец не согласился на обоюдную отмену заказа

---

## deleteCancelOrderRequestByWorker

**Path:** `POST /deleteCancelOrderRequestByWorker`

**Summary:** Продавец удалил свой запрос на обоюдную отмену заказа

---

## workerInprogress

**Path:** `POST /workerInprogress`

**Summary:** Продавец взял заказ в работу

---

## workerDeclineExtras

**Path:** `POST /workerDeclineExtras`

**Summary:** Продавец отклоняет предложенные опции

---

## getCompanyDetails

**Path:** `POST /getCompanyDetails`

**Summary:** Получение деталей о компании по ИНН

---

## sendCompanyForVerification

**Path:** `POST /sendCompanyForVerification`

**Summary:** Отправка компании на проверку

---

