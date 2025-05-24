import type { FieldPath, FieldValues } from "react-hook-form";
import SelectFieldEnergyForm from "../../../shared/SelectFieldEnergyForm/UI/SelectFieldEnergyForm";
import TextFieldEnergyForm from "../../../shared/TextFieldEnergyForm/UI/TextFieldEnergyForm";
import {
	fileFields,
	monthsTextFields,
	selectFields,
	textFields,
	type FieldsEnergyFormProps,
} from "../config/fieldsEnergyFormConfig";
import styles from "../styles/FieldsEnergyForm.module.css";
import { useState } from "react";
import FileFieldEnergyForm from "../../../shared/FileFieldEnergyForm/UI/FileFieldEnergyForm";

const FieldsEnergyForm = <T extends FieldValues>({
	register,
}: FieldsEnergyFormProps<T>) => {
	const [isMonthsVisible, setMonthsVisible] = useState(false);
	const handleChangeVisible = () => setMonthsVisible((prev) => !prev);

	return (
		<div className={styles.mainBlock}>
			{textFields.map(({ title, field }) => (
				<TextFieldEnergyForm
					field={field as FieldPath<T>}
					register={register}
					title={title}
				/>
			))}
			{selectFields.map(({ title, field, options }) => (
				<SelectFieldEnergyForm
					register={register}
					field={field as FieldPath<T>}
					title={title}
					options={options}
				/>
			))}
			<div>
				<label onClick={handleChangeVisible} className={styles.toggleLabel}>
					Потребление энергии (в квт/ч)
					<span className={styles.toggleIcon}>
						{isMonthsVisible ? "▼" : "►"}
					</span>
				</label>
				<div className={isMonthsVisible ? styles.monthsBlock : ""}>
					{isMonthsVisible &&
						monthsTextFields.map(({ month, field }) => (
							<TextFieldEnergyForm
								register={register}
								field={field as FieldPath<T>}
								title={month}
							/>
						))}
				</div>
			</div>
			{fileFields.map(({ field, title }) => (
				<FileFieldEnergyForm
					title={title}
					field={field as FieldPath<T>}
					register={register}
				/>
			))}
		</div>
	);
};

export default FieldsEnergyForm;
