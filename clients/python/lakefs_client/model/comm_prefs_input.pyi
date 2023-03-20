# coding: utf-8

"""
    lakeFS API

    lakeFS HTTP API  # noqa: E501

    The version of the OpenAPI document: 0.1.0
    Contact: services@treeverse.io
    Generated by: https://openapi-generator.tech
"""

from datetime import date, datetime  # noqa: F401
import decimal  # noqa: F401
import functools  # noqa: F401
import io  # noqa: F401
import re  # noqa: F401
import typing  # noqa: F401
import typing_extensions  # noqa: F401
import uuid  # noqa: F401

import frozendict  # noqa: F401

from lakefs_client import schemas  # noqa: F401


class CommPrefsInput(
    schemas.DictSchema
):
    """NOTE: This class is auto generated by OpenAPI Generator.
    Ref: https://openapi-generator.tech

    Do not edit the class manually.
    """


    class MetaOapg:
        required = {
            "featureUpdates",
            "securityUpdates",
        }
        
        class properties:
            featureUpdates = schemas.BoolSchema
            securityUpdates = schemas.BoolSchema
            email = schemas.StrSchema
            __annotations__ = {
                "featureUpdates": featureUpdates,
                "securityUpdates": securityUpdates,
                "email": email,
            }
    
    featureUpdates: MetaOapg.properties.featureUpdates
    securityUpdates: MetaOapg.properties.securityUpdates
    
    @typing.overload
    def __getitem__(self, name: typing_extensions.Literal["featureUpdates"]) -> MetaOapg.properties.featureUpdates: ...
    
    @typing.overload
    def __getitem__(self, name: typing_extensions.Literal["securityUpdates"]) -> MetaOapg.properties.securityUpdates: ...
    
    @typing.overload
    def __getitem__(self, name: typing_extensions.Literal["email"]) -> MetaOapg.properties.email: ...
    
    @typing.overload
    def __getitem__(self, name: str) -> schemas.UnsetAnyTypeSchema: ...
    
    def __getitem__(self, name: typing.Union[typing_extensions.Literal["featureUpdates", "securityUpdates", "email", ], str]):
        # dict_instance[name] accessor
        return super().__getitem__(name)
    
    
    @typing.overload
    def get_item_oapg(self, name: typing_extensions.Literal["featureUpdates"]) -> MetaOapg.properties.featureUpdates: ...
    
    @typing.overload
    def get_item_oapg(self, name: typing_extensions.Literal["securityUpdates"]) -> MetaOapg.properties.securityUpdates: ...
    
    @typing.overload
    def get_item_oapg(self, name: typing_extensions.Literal["email"]) -> typing.Union[MetaOapg.properties.email, schemas.Unset]: ...
    
    @typing.overload
    def get_item_oapg(self, name: str) -> typing.Union[schemas.UnsetAnyTypeSchema, schemas.Unset]: ...
    
    def get_item_oapg(self, name: typing.Union[typing_extensions.Literal["featureUpdates", "securityUpdates", "email", ], str]):
        return super().get_item_oapg(name)
    

    def __new__(
        cls,
        *_args: typing.Union[dict, frozendict.frozendict, ],
        featureUpdates: typing.Union[MetaOapg.properties.featureUpdates, bool, ],
        securityUpdates: typing.Union[MetaOapg.properties.securityUpdates, bool, ],
        email: typing.Union[MetaOapg.properties.email, str, schemas.Unset] = schemas.unset,
        _configuration: typing.Optional[schemas.Configuration] = None,
        **kwargs: typing.Union[schemas.AnyTypeSchema, dict, frozendict.frozendict, str, date, datetime, uuid.UUID, int, float, decimal.Decimal, None, list, tuple, bytes],
    ) -> 'CommPrefsInput':
        return super().__new__(
            cls,
            *_args,
            featureUpdates=featureUpdates,
            securityUpdates=securityUpdates,
            email=email,
            _configuration=_configuration,
            **kwargs,
        )