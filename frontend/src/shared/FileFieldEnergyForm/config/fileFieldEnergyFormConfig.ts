import type { FieldPath, FieldValues, UseFormRegister } from "react-hook-form";

export interface IFileFieldEnergyForm<T extends FieldValues> {
	register: UseFormRegister<T>;
	field: FieldPath<T>;
	title: string;
}
