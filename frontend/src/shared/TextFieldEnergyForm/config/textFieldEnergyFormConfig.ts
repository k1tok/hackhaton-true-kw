import type { FieldPath, FieldValues, UseFormRegister } from "react-hook-form";

export interface ITextFieldEnergyForm<T extends FieldValues> {
	register: UseFormRegister<T>;
	title: string;
	field: FieldPath<T>;
}
