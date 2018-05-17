qskr
====

golang euc-kr url encode / decode


utf8 아닌 예전 euc-kr로 구성된 사이트를 스크랩할때 웹서버로 보내는 데이터를 euc-kr로 인코딩해주는 함수. 


사용법
------

```
enc := URLEncode(in)

dec,err := URLDecode(enc)
```

샘플코드는 test code 참조.
