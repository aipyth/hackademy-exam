import React, {useState} from "react";
import requests from "../utils/requests";

export default function(): JSX.Element {
    const [inputActive, setInputActive] = useState(false)
    const [newListName, setNewListName] = useState('')

    const onNewListChange = (event: React.ChangeEvent<HTMLInputElement>) => {
        setNewListName(event.target.value)
    }

    const addList = () => {
        requests.createList({
            "name": newListName,
        }).then(res => {
            console.log('created list', res)
        })
    }

    const classes = {
        'active': "w-full text-lg taxt-gray-700 bg-white flex flex-row cursor-pointer",
        'unactive': "w-full text-lg text-black flex flex-row",
    }

    const renderUnactive = (): JSX.Element => {
        return (
            <p>New list</p>
        )
    }
    const renderActive = (): JSX.Element => {
        return (
            <div className="flex flex-row">
                <input
                    className="w-2/3"
                    type="text"
                    onChange={onNewListChange}
                    placeholder="New list"
                />
                <span
                    className="text-yellow-300 w-1/3 px-1"
                    onClick={addList}>Add</span>
            </div>
        )
    }

    return (
        <div 
            className={ inputActive ? classes['active'] : classes['unactive'] } 
            onClick={() => setInputActive(true)}
        >

            <span className="mx-2 text-xl cursor-pointer"
                onClick={() => { setInputActive(!inputActive) }}>
                +
            </span>
            { inputActive ? renderActive() : renderUnactive() }
        </div>
    )
}
