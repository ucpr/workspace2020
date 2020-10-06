import { FunctionalComponent, h } from "preact";
import styled from "styled-components";

import Header from "./header";

const App: FunctionalComponent = () => {
    // left content
    const DIVA = styled.div`
        min-height: 100vh;
        display: flex;
        flex-wrap: wrap;
        align-items: flex-start;
    `;

    // right content
    const DIVB = styled.div`
        display: flex;
        flex-direction: column;
 `;

    return (
        <div id="app">
            <DIVA>
              <a>content A</a>
            </DIVA>
            <DIVB>
                <div>content B</div>
                <div>content C</div>
            </DIVB>
        </div>
    );
};

export default App;
