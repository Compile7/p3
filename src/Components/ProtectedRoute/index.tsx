import React from 'react';
import { Navigate, Route, RouteProps, useLocation } from 'react-router-dom';
import { isLoggedIn } from '../../utils/index';

export const ProtectedRoute: React.FC<RouteProps> = (props) => {
    const { pathname, search } = useLocation();

    if (isLoggedIn()) {
        const { children, ...otherProps } = props;
        return <Route {...otherProps}>{children as any}</Route>;
    } else
        return search ? (
            <Navigate replace to={`/signin?redirect=${pathname}&${search.slice(1)} `} />
        ) : (
            <Navigate replace to={`/signin?redirect=${pathname}`} />
        );
};
