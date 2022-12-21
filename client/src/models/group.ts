export default class Group {
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

	constructor(id: Number, description: string) {
		this._id = id
		this._description = description
	}
}
