# #360开源模型重塑AI视觉理解##360解决CLIP近视问题#CLIP的“近视”问题，被360搞定了。360人工智能研究院最新图文跨模态模型FG-CLIP，宣布以“长文本深度理解”...

**URL**: https://weibo.com/6105753431/PpwhSDNRJ

## 原始摘要

<a href="https://m.weibo.cn/search?containerid=231522type%3D1%26t%3D10%26q%3D%23360%E5%BC%80%E6%BA%90%E6%A8%A1%E5%9E%8B%E9%87%8D%E5%A1%91AI%E8%A7%86%E8%A7%89%E7%90%86%E8%A7%A3%23&amp;extparam=%23360%E5%BC%80%E6%BA%90%E6%A8%A1%E5%9E%8B%E9%87%8D%E5%A1%91AI%E8%A7%86%E8%A7%89%E7%90%86%E8%A7%A3%23" data-hide=""><span class="surl-text">#360开源模型重塑AI视觉理解#</span></a><a href="https://m.weibo.cn/search?containerid=231522type%3D1%26t%3D10%26q%3D%23360%E8%A7%A3%E5%86%B3CLIP%E8%BF%91%E8%A7%86%E9%97%AE%E9%A2%98%23&amp;extparam=%23360%E8%A7%A3%E5%86%B3CLIP%E8%BF%91%E8%A7%86%E9%97%AE%E9%A2%98%23" data-hide=""><span class="surl-text">#360解决CLIP近视问题#</span></a><br><br>CLIP的“近视”问题，被360搞定了。<br><br>360人工智能研究院最新图文跨模态模型FG-CLIP，宣布以“长文本深度理解”和“细粒度视觉比对”双突破，彻底解决了传统CLIP模型的“视觉近视”问题，能够精准识别局部细节。<br><br>具体怎么个说法？先来个视力大挑战：找一找图中的哪一句话，正确描述了图像左边的内容？【图1】<br><br>正确答案是：“A light brown wood stool（一个浅棕色的木凳子）”，注意看，这个木凳子位于画面的中央偏右，悄悄隐藏在狗狗的身后。【图2】<br><br>可以发现，4个常用模型——CLIP、EVACLIP、SIGLIP、FINE-CLIP基于左侧图片选出的最匹配的文本描述是：A blue dog with a white colored head。<br><br>显然这个描述是错误的，这就是CLIP的“视觉近视”问题：会因为对比损失倾向于拉近全局图像与文本的嵌入，而非局部区域的对齐，削弱了细粒度特征学习。<br><br>而FG-CLIP则精准命中了答案。<br><br>实验结果显示，FG-CLIP在细粒度理解、开放词汇对象检测、长短文本图文检索以及通用多模态基准测试等下游任务中均显著优于原始CLIP和其他最先进方法。【图3】<br><br>在12个下游任务上，FG-CLIP相比现有模型在关键的长文本理解+细粒度比对上实现了大幅突破。<br><br>具体是怎么做到的？欢迎点击链接查看更多技术细节：<a href="https://weibo.cn/sinaurl?u=https%3A%2F%2Fmp.weixin.qq.com%2Fs%2Fmw02qHnJsZ13A1adeZ6Nxw" data-hide=""><span class="url-icon"><img style="width: 1rem;height: 1rem" src="https://h5.sinaimg.cn/upload/2015/09/25/3/timeline_card_small_web_default.png" referrerpolicy="no-referrer"></span><span class="surl-text">告别“图文不符”！FG-CLIP实现细粒度跨模态对齐，360开源模型重塑AI视觉理解</span></a><br>360人工智能研究院还表示，将全面开源模型及其相关数据：<br><a href="https://weibo.cn/sinaurl?u=https%3A%2F%2Fgithub.com%2F360CVGroup" data-hide=""><span class="url-icon"><img style="width: 1rem;height: 1rem" src="https://h5.sinaimg.cn/upload/2015/09/25/3/timeline_card_small_web_default.png" referrerpolicy="no-referrer"></span><span class="surl-text">网页链接</span></a><img style="" src="https://tvax1.sinaimg.cn/large/006Fd7o3gy1i0xtgicvduj30u00ep11x.jpg" referrerpolicy="no-referrer"><br><br><img style="" src="https://tvax1.sinaimg.cn/large/006Fd7o3gy1i0xtgkbw5hj30u00b3tf8.jpg" referrerpolicy="no-referrer"><br><br><img style="" src="https://tvax2.sinaimg.cn/large/006Fd7o3gy1i0xtgnxgznj30u00p6n8q.jpg" referrerpolicy="no-referrer"><br><br>

## AI 摘要

360人工智能研究院开源了FG-CLIP模型，通过"长文本深度理解"和"细粒度视觉比对"技术，解决了传统CLIP模型的"视觉近视"问题。该模型能精准识别图像局部细节，在12个下游任务测试中表现优于CLIP等现有模型。实验显示FG-CLIP能正确识别被遮挡物体（如隐藏在狗后的木凳），而其他模型会错误匹配全局特征。该突破实现了细粒度跨模态对齐，显著提升了图文检索、对象检测等任务的准确性。360宣布将全面开源模型及相关数据。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-04-29T19:02:47Z
- **目录日期**: 2025-04-29
