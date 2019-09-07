package domain

import "time"

type InventoryProduct struct {
	Barcode         string  `json:"barcode,omitempty"`          //  штрих-код
	Name            string  `json:"name,omitempty"`             // наименование
	ActualAmount    float64 `json:"actual_amount,omitempty"`    // фактическое количество
	ResidueAmount   float64 `json:"residue_amount,omitempty"`   // остаток количества
	DiffAmount      float64 `json:"diff_amount,omitempty"`      // разница количества
	CostPrice       float64 `json:"cost_price,omitempty"`       // себестоимость
	ActualMoney     float64 `json:"actual_money,omitempty"`     // фактическая сумма
	AccountingMoney float64 `json:"accounting_money,omitempty"` // учетная сумма
	DiffMoney       float64 `json:"diff_money"`                 // разница суммы
	MeasureUnit     string  `json:"measure_unit,omitempty"`     // единица измерения
	NomenclatureId  string  `json:"nomenclature_id,omitempty"`  // продукт ID
}

type Inventory struct {
	Products                 []InventoryProduct `json:"products,omitempty"`                    // массив продуктов
	Comment                  string             `json:"comment,omitempty"`                     // коментарий
	SellDuringInventory      bool               `json:"sell_during_inventory,omitempty"`       // учитывать продажи во время ревизии
	ResetNotScanningProducts bool               `json:"reset_not_scanning_products,omitempty"` // обнулить количество непросканированных товаров
	ShortInventory                              // информация о резии
}

type ShortInventory struct {
	InventoryId     *int      `json:"inventory_id,omitempty"`     // номер ревизии
	EmployeeId      string    `json:"employee_id,omitempty"`      // ID сотрудника
	StockId         string    `json:"stock_id,omitempty"`         // ID склада
	StockName       string    `json:"stock_name,omitempty"`       // название склада
	AccountingMoney float64   `json:"accounting_money,omitempty"` // учетная сумма
	ActualMoney     float64   `json:"actual_money,omitempty"`     // фактическая сумма
	DiffMoney       float64   `json:"diff_money"`                 // разница суммы
	CreatedOn       time.Time `json:"created_on,omitempty"`       // дата создания
	UpdatedOn       time.Time `json:"updated_on,omitempty"`       // дата обновления
	Status          string    `json:"status,omitempty"`           // статус ревизии
	MerchantId      string    `json:"merchant_id,omitempty"`      // ID продавца
	Id              string    `json:"id,omitempty"`               // ID в нашей системе
	Comment         string    `json:"comment,omitempty"`
}
