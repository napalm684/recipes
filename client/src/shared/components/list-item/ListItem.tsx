import * as React from 'react';

import './ListItem.css';

export interface IListItemProps {
    label: string
}

export default function ListItem(props: IListItemProps){
    return (
        <section className="ListItem">
            {props.label}
        </section>  
    );
}
