import type { FieldValues, UseFormRegister } from "react-hook-form";

export interface FieldsEnergyFormProps<T extends FieldValues> {
	register: UseFormRegister<T>;
}

export const textFields = [
	{ title: "Номер обращения", field: "accountId" },
	{
		title: "Адрес",
		field: "address",
	},
	{
		title: "Количество комнат",
		field: "roomsCount",
	},
	{
		title: "Количество жильцов",
		field: "residentsCount",
	},
	{
		title: "Площадь помещения (в м²)",
		field: "totalArea",
	},
];

export const selectFields = [
	{
		title: "Тип строения",
		field: "buildingType",
		options: ["Частный", "Другое"],
	},
];

export const monthsTextFields = [
	{ month: "Январь", field: `consumption.${1}` },
	{ month: "Февраль", field: `consumption.${2}` },
	{ month: "Март", field: `consumption.${3}` },
	{ month: "Апрель", field: `consumption.${4}` },
	{ month: "Май", field: `consumption.${5}` },
	{ month: "Июнь", field: `consumption.${6}` },
	{ month: "Июль", field: `consumption.${7}` },
	{ month: "Август", field: `consumption.${8}` },
	{ month: "Сентябрь", field: `consumption.${9}` },
	{ month: "Октябрь", field: `consumption.${10}` },
	{ month: "Ноябрь", field: `consumption.${11}` },
	{ month: "Декабрь", field: `consumption.${12}` },
];

export const fileFields = [{ title: "Загрузить файл", field: "file" }];
