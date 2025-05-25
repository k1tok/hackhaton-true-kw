import type { FieldPath, FieldValues, UseFormRegister } from "react-hook-form";

export interface ISelectFieldEnergyForm<T extends FieldValues> {
	register: UseFormRegister<T>;
	title: string;
	field: FieldPath<T>;
	options: string[];
}
