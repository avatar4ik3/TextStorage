export default class Text {
	private _id: Number
	public get id(): Number {
		return this._id
	}
	public set id(v: Number) {
		this._id = v
	}

	private _description: string
	public get description(): string {
		return this._description
	}
	public set description(v: string) {
		this._description = v
	}

	private _value: string
	public get value(): string {
		return this._value
	}
	public set value(v: string) {
		this._value = v
	}

	constructor(id: Number, description: string, value: string) {
		this._id = id
		this._value = value
		this._description = description
	}
}
