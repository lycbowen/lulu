import { Carousel } from "antd"
import './App.css'
import { GetPicFitMode } from '../wailsjs/go/main/App'
import { useEffect, useState } from "react"
import React from "react"

const images = Object.values(
    import.meta.glob('./assets/images/*.{jpg,jpeg,png}', {
        eager: true,
        import: 'default',
    })
) as string[]

function App() {
    const [fitMode, setFitMode] = useState<React.CSSProperties['objectFit']>('cover');

    useEffect(() => {
        GetPicFitMode().then((res) => setFitMode(res as React.CSSProperties['objectFit']))
    }, []);

    return (
        <Carousel className="container" fade speed={1000} autoplay autoplaySpeed={5000} draggable>
            {images.map((img, i) => (
                <div className="slide" key={i}>
                    <img className="pic" src={img} style={{ objectFit: fitMode }} />
                </div>
            ))}
        </Carousel>
    )
}

export default App
