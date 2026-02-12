# Описание предметной области: Банк

## Цель
Моделирование домена банковской системы с использованием принципов ООП и расширение связей между сущностями с помощью наследования, композиции и агрегации. Проект описывает базовые операции: открытие счетов, пополнение, снятие, перевод, фиксацию транзакций и публикацию событий.

## Основные сущности

### Value Object
`Money` описывает денежную величину и валюту. Поля закрыты, изменение только через операции сложения/вычитания.

```
type Money struct {
    amount int
    curr   Currency
}
type Currency string
```

### Перечисления

```
type TransactionType int
const (
    Transfer TransactionType = iota + 1
    Deposit
    Withdraw
)
```

```
type PaymentStatus int
const (
    Initiated PaymentStatus = iota
    Pending
    Declined
    Failed
    Completed
)
```

```
type AccountStatus int
const (
    Active AccountStatus = iota + 1
    Frozen
    Closed
)
```

```
type BonusTier int
const (
    Bronze BonusTier = iota + 1
    Silver
    Gold
)
```

### Entity
`Client` — клиент банка. Содержит персональные данные и список счетов (агрегация).

```
Client
- clientUUID: uuid.UUID
- Person: Person
- accounts: []AccountRef
```

`Person` — персональные данные (композиция в составе `Client`).

```
Person
- firstname: string
- lastname: string
- surname: string
- birthdayDate: time.Time
```

`Account` — базовый счет. Используется как основа для специализированных типов.

```
Account
- userUUID: uuid.UUID
- accountUUID: uuid.UUID
- status: AccountStatus
- balance: Money
```

#### Виды счетов
1. **CheckingAccount** — базовый расчетный счет.
2. **SavingsAccount** — накопительный счет с бонусной программой.
3. **BonusAccount** — расширенный накопительный счет с дополнительным множителем бонусов.
4. **CreditAccount** — кредитный счет с `overdraftLimit`.

`Transaction` — операция перевода или списания/пополнения.

```
Transaction
- id: int
- srcAccountUUID: uuid.UUID
- destAccountUUID: uuid.UUID
- value: Money
- status: PaymentStatus
- type: TransactionType
```

`Event` — доменное событие (например, успешный перевод).

```
Event
- Name: string
- AccountUUID: uuid.UUID
- OccurredAt: time.Time
```

## Services & Interfaces
`PaymentAccount` — интерфейс для работы со счетами в прикладном слое.

```
type PaymentAccount interface {
    UserUUID() uuid.UUID
    AccountUUID() uuid.UUID
    Status() AccountStatus
    Balance() Money
    Deposit(m Money)
    Withdraw(m Money) bool
    CanWithdraw(m Money) bool
    SetStatus(s AccountStatus)
}
```

`AccountRepository` — абстракция хранилища счетов.

```
type AccountRepository interface {
    Create(account PaymentAccount) error
    ChangeStatus(accountUUID uuid.UUID, status AccountStatus) error
    ByUUID(accountUUID uuid.UUID) (PaymentAccount, error)
}
```

`EventService` — публикация событий.

```
type EventService interface {
    Publish(event Event) error
    QueryAll() []Event
}
```

`PaymentService` — сервис пополнения/снятия/переводов.

```
type PaymentService interface {
    Deposit(m Money, aUUID uuid.UUID) error
    Withdraw(m Money, aUUID uuid.UUID) error
    ProcessTransaction(t Transaction) error
}
```

## Use Case
`TransferUseCase` оркестрирует перевод: определяет комиссию по типу счета, выполняет списание/зачисление, обновляет статус транзакции и публикует события.

## Бизнес-правила
1. Снятие возможно только если `CanWithdraw` возвращает `true`.
2. При статусе счета `Frozen` или `Closed` операции пополнения и списания не выполняются.
3. Для кредитного счета учитывается `overdraftLimit`.
4. Для накопительного/бонусного счета при пополнении начисляются бонусные баллы.
5. Валюта операции должна совпадать с валютой счета.

## Связь с принципами ООП
- **Инкапсуляция:** закрытые поля, доступ через методы.
- **Наследование:** специализированные счета расширяют базовый `Account`.
- **Композиция:** `Account` содержит `Money`, `SavingsAccount` содержит `BonusProgram`, `Client` содержит `Person`.
- **Агрегация:** `Client` хранит ссылки на счета через `AccountRef`.
- **Ассоциация:** `Transaction` знает о счетах через их UUID.
- **Полиморфизм:**
  - Подтипный: `PaymentAccount` и разные типы счетов.
  - Ad hoc: расчет комиссии через `CalculateFee` с `type switch`.
  - Параметрический: обобщенные `InMemoryMap`/`InMemoryList`.
