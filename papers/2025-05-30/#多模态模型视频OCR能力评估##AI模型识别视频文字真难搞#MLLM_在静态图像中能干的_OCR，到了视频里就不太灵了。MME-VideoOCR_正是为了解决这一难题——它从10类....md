# #多模态模型视频OCR能力评估##AI模型识别视频文字真难搞#MLLM 在静态图像中能干的 OCR，到了视频里就不太灵了。MME-VideoOCR 正是为了解决这一难题——它从10类...

**URL**: https://weibo.com/6105753431/Pue1omeMh

## 原始摘要

<a href="https://m.weibo.cn/search?containerid=231522type%3D1%26t%3D10%26q%3D%23%E5%A4%9A%E6%A8%A1%E6%80%81%E6%A8%A1%E5%9E%8B%E8%A7%86%E9%A2%91OCR%E8%83%BD%E5%8A%9B%E8%AF%84%E4%BC%B0%23&amp;extparam=%23%E5%A4%9A%E6%A8%A1%E6%80%81%E6%A8%A1%E5%9E%8B%E8%A7%86%E9%A2%91OCR%E8%83%BD%E5%8A%9B%E8%AF%84%E4%BC%B0%23" data-hide=""><span class="surl-text">#多模态模型视频OCR能力评估#</span></a><a href="https://m.weibo.cn/search?containerid=231522type%3D1%26t%3D10%26q%3D%23AI%E6%A8%A1%E5%9E%8B%E8%AF%86%E5%88%AB%E8%A7%86%E9%A2%91%E6%96%87%E5%AD%97%E7%9C%9F%E9%9A%BE%E6%90%9E%23&amp;extparam=%23AI%E6%A8%A1%E5%9E%8B%E8%AF%86%E5%88%AB%E8%A7%86%E9%A2%91%E6%96%87%E5%AD%97%E7%9C%9F%E9%9A%BE%E6%90%9E%23" data-hide=""><span class="surl-text">#AI模型识别视频文字真难搞#</span></a><br><br>MLLM 在静态图像中能干的 OCR，到了视频里就不太灵了。MME-VideoOCR 正是为了解决这一难题——它从10类任务、25个子任务全方位评估模型能否“看清”“理解”“推理”视频中的文字。<br><br>数据集方面，MME-VideoOCR共包含1,464个视频片段、2,000条问答对，覆盖日常生活、科普、游戏、AIGC等多样场景，还专门加了模糊、遮挡、低清等干扰项，挑战性拉满。<br><br>任务上不仅考验基础识别，还要求时空定位、时序追踪、特殊文本解析、信息整合与逻辑推理。评估方式也很灵活，用了字符串匹配、多选题和GPT辅助评分三板斧。<br><br>结果揭示现实之“骨感”：就算是最强模型Gemini-2.5 Pro，也只有73.7%的准确率。开源模型甚至连60%都没摸到，主要卡在“时序理解”和“文字推理”这两关。 <a href="https://weibo.com/ttarticle/p/show?id=2309405172037424644248" data-hide=""><span class="url-icon"><img style="width: 1rem;height: 1rem" src="https://h5.sinaimg.cn/upload/2015/09/25/3/timeline_card_small_article_default.png" referrerpolicy="no-referrer"></span><span class="surl-text">全面评估多模态模型视频OCR能力，Gemini准确率仅73.7%</span></a><img style="" src="https://tvax3.sinaimg.cn/large/006Fd7o3gy1i1xlehaif0j30qk0ey76y.jpg" referrerpolicy="no-referrer"><br><br>

## AI 摘要

MME-VideoOCR是一个专门评估多模态模型视频OCR能力的新基准，包含1,464个视频片段和2,000条问答对，覆盖10类任务25个子任务。该基准不仅测试基础文字识别，还挑战模型的时空定位、时序追踪和逻辑推理能力，并加入了模糊、遮挡等干扰项。评估结果显示，即使是表现最好的Gemini-2.5 Pro模型准确率也仅为73.7%，开源模型普遍低于60%，主要困难在于时序理解和文字推理。这表明当前AI在视频文字识别和处理方面仍面临重大挑战。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-05-30T16:03:45Z
- **目录日期**: 2025-05-30
