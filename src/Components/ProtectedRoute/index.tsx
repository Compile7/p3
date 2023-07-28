import React from 'react';
import { Navigate, RouteProps, useLocation } from 'react-router-dom';
import { isLoggedIn } from '../../utils/index';
import { useStateContext } from '../../Contexts/contextProvider';

type ProtectedRouteProps = RouteProps & {
    roleProtected?: boolean
}
export const ProtectedRoute: React.FC<ProtectedRouteProps> = (props) => {
    const { pathname, search } = useLocation();
    const { children } = props;
    const { profile } = useStateContext();
    console.log(profile)
    if (!isLoggedIn()) {

        return search ? (
            <Navigate to={`/?redirect=${pathname}&${search.slice(1)} `} replace />

        ) : (
            <Navigate to={`/?redirect=${pathname}`} replace />

        );
    } else if (props.roleProtected) {

        return profile?.Role === 'Admin' ? children : search ? (
            <Navigate to={`/?redirect=${pathname}&${search.slice(1)} `} replace />

        ) : (
            <Navigate to={`/?redirect=${pathname}`} replace />

        );
    } else return children
};
