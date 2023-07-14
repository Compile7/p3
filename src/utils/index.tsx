export const isLoggedIn=()=>!!localStorage.getItem("P3AccessToken")

export const FALLBACK = (
    <div id='preloader' className='fullscreen'>
        <div id='status'>
            <div className='spinner'>
                <div className='double-bounce1'></div>
                <div className='double-bounce2'></div>
            </div>
        </div>
    </div>
);