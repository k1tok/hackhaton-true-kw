import type { FieldValues } from "react-hook-form";
import type { ISelectFieldEnergyForm } from "../config/selectFieldEnergyFormConfig";
import styles from "../style/SelectFieldEnergyForm.module.css";

const SelectFieldEnergyForm = <T extends FieldValues>({
	title,
	register,
	field,
	options,
}: ISelectFieldEnergyForm<T>) => {
	return (
		<div className={styles.mainBlock}>
			<label className={styles.title}>{title}</label>
			<select {...register(field)} className={styles.select}>
				<option value="" disabled selected>
					Выберите
				</option>
				{options.map((opt) => (
					<option value={opt}>{opt}</option>
				))}
			</select>
		</div>
	);
};

export default SelectFieldEnergyForm;
