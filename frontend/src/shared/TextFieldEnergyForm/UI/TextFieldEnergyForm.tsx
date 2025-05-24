import type { FieldValues } from "react-hook-form";
import type { ITextFieldEnergyForm } from "../config/textFieldEnergyFormConfig";
import styles from "../styles/TextFieldEnergyForm.module.css";

const TextFieldEnergyForm = <T extends FieldValues>({
	register,
	title,
	field,
}: ITextFieldEnergyForm<T>) => {
	return (
		<div className={styles.block}>
			<label className={styles.title}>{title}</label>
			<input className={styles.input} type="text" {...register(field)} />
		</div>
	);
};

export default TextFieldEnergyForm;
