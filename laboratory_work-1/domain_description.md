# Описание предметной области: Банк

## Цель
Моделирование домена(базовые сущности и процессы) - банковской системы с использованием принципов ООП. Проект в себе несет базовые операции без использования излишней бизнес логики : открытие счетов, пополнение, снятие, перевод, фиксацию транзакций и публикацию событий.

## Основные сущности

### Value Object: 

`Money` описывает денежную величину и валюту. Это неизменяемый объект, используемый везде, где требуется представление сумм

```
type Money struct {
    amount   int64
    currency Currency
}
type Currency string
```

`type value_object_name type` - value-object значения и перечисления базовых состояний процессов
```
type TransactionType int - тип банковской транзакции, для корректной обработки и разделения типов транзакций
const (
    Transfer TransactionType = iota + 1
    Deposit
    Withdraw
)
```

```
type PaymentStatus int - текущий статус платежа, для изменения состояние тразакции и создание ивентов 
const (
    Initiated PaymentStatus = iota + 1
    Pending
    Declined
    Failed
    Completed
)
```

```
type AccountStatus int - статус банковского акканута, для отслеживая возможности проведения операций и поддержания логики
const (
    Active AccountStatus = iota + 1
    Frozen
    Closed
)
```

### Entity:
`Client` Клиент банка. Содержит персональные данные и идентификатор. Выделен интерфейс для инкапсуляции полей класса и добавления валидаций, обеспечения консистености данных

```
Client
- id: uuid
- firstName: string
- lastName: string
- surName: string
- birthdayDate: time.Time
```

`Account` Счет принадлежит клиенту, имеет уникальный идентификатор, статус и баланс. Используется как базовая сущность для кастомных счетов. Выделен интерфейс для инкапсуляции полей класса и добавления валидаций, обеспечения консистености данных

```
Account
- userUUID: uuid
- status: AccountStatus
- balance: Money
```

#### Виды счетов 
1. **Накопительный (SavingsAccount)**  
   Имеет бонусную программу; при пополнении начисляет бонусные баллы
2. **Кредитный (CreditAccount)**  
   Имеет `overdraftLimit` и разрешает уходить в минус в пределах лимита
3. **Бонусный (BonusAccount)**  
   Savings с расширенной программой лояльности.

`Transaction` Описывает перевод или операцию списания/пополнения. Выделен интерфейс для инкапсуляции полей класса и добавления валидаций, обеспечения консистености данных

```
Transaction
- id: int
- sourceUUID: uuid
- destinationUUID: uuid
- amount: Money
- status: PaymentStatus
- type: TransactionType
```

## Services & Interfaces

`PaymentAccount` Интерфейс для работы со счетами в прикладном слое

```
type PaymentAccount interface {
    UserUUID() uuid.UUID
    Deposit(value Money)
    Withdraw(value Money) bool
    CanWithdraw(value Money) bool
}
```

`AccountRepository` Абстракция хранилища счетов

```
type AccountRepository interface {
    Create(account PaymentAccount) error
    ChangeStatus(accountUUID uuid.UUID, status AccountStatus) error
}
```

`EventService`Сервис публикации событий (логи, метрики, etc)

```
type EventPublisher interface {
    Publish(event Event)
}
```

`ProcessPaymentService` Сервис обработки платежей:
```
type PaymentService interface {
	Deposit(m domain.Money, aUUID uuid.UUID) error
	Withdraw(m domain.Money, aUUID uuid.UUID) error
	ProcessTransaction(t domain.Transaction) error
}
```

## Бизнес-правила 
1. Снятие возможно только если `CanWithdraw` возвращает `true`.
2. При статусе счета `Frozen` или `Closed` операции пополнения и списания не выполняются.
3. Для кредитного счета учитывается `overdraftLimit`.
4. Для накопительного/бонусного счета при пополнении начисляются бонусные баллы.

## Связь с принципами ООП
- **Инкапсуляция:** поля сущностей закрыты, доступ через методы.
- **Абстракция:** через интерфейсы `PaymentAccount`, `AccountRepository`, `EventPublisher`.
- **Полиморфизм:** разные типы счетов реализуют общее поведение по-разному.
- **Наследование/композиция:** использование встраивания базового `Account` в специализированные типы.
