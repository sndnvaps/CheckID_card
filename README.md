#用于检验身份证的正确性，主要针对第二代身份证（18位）
 
[![Build Status](https://travis-ci.org/sndnvaps/CheckID_card.svg?branch=master)](http://travis-ci.org/sndnvaps/CheckID_card)




CheckID_card 用来检验身份证的正确性

计算方法如下:

 	ai -> a1 , a2, a3, a4, a5, a6... a17 (a18 是校验码) 身份证前17位对应(ai)
	wi -> 7, 9, 10, 5, 8, 4, 2, 1, 6, 3, 7, 9, 10, 5, 8, 4, 2 (17位)

	int  res = 0;
	for (i = 1; i < 17; i++) {
	 res += (a[i] * w[i])
	}
	int  y = res % 11；


  y 与 a18的对应关系

	y    { 0, 1,  2,  3, 4, 5, 6, 7, 8, 9, 10}
	a18  { 1, 0, 'X', 9, 8, 7, 6, 5, 4, 3,  2 } -> vefiry[18] = { 1, 0, 'X', 9, 8, 7, 6, 5, 4, ,3, 2};


 运算效果：
	
	CheckID_card
	34052419800101001X
 	身份证号码是 =  34052419800101001X
	身份证最后一位是  X
 	true 验证成功
 	
 	CheckID_card
	340524198001010013
 	身份证号码是 =  340524198001010013
	计算得到身份证最后一位是  X
 	身份证最后一位是  3
 	false 验证失败
 
 
 





