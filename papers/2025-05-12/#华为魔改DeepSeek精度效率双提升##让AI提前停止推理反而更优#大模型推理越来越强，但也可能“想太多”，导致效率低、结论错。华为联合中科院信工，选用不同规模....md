# #华为魔改DeepSeek精度效率双提升##让AI提前停止推理反而更优#大模型推理越来越强，但也可能“想太多”，导致效率低、结论错。华为联合中科院信工，选用不同规模...

**URL**: https://weibo.com/6105753431/PrnmkDtMy

## 原始摘要

<a href="https://m.weibo.cn/search?containerid=231522type%3D1%26t%3D10%26q%3D%23%E5%8D%8E%E4%B8%BA%E9%AD%94%E6%94%B9DeepSeek%E7%B2%BE%E5%BA%A6%E6%95%88%E7%8E%87%E5%8F%8C%E6%8F%90%E5%8D%87%23&amp;extparam=%23%E5%8D%8E%E4%B8%BA%E9%AD%94%E6%94%B9DeepSeek%E7%B2%BE%E5%BA%A6%E6%95%88%E7%8E%87%E5%8F%8C%E6%8F%90%E5%8D%87%23" data-hide=""><span class="surl-text">#华为魔改DeepSeek精度效率双提升#</span></a><a href="https://m.weibo.cn/search?containerid=231522type%3D1%26t%3D10%26q%3D%23%E8%AE%A9AI%E6%8F%90%E5%89%8D%E5%81%9C%E6%AD%A2%E6%8E%A8%E7%90%86%E5%8F%8D%E8%80%8C%E6%9B%B4%E4%BC%98%23&amp;extparam=%23%E8%AE%A9AI%E6%8F%90%E5%89%8D%E5%81%9C%E6%AD%A2%E6%8E%A8%E7%90%86%E5%8F%8D%E8%80%8C%E6%9B%B4%E4%BC%98%23" data-hide=""><span class="surl-text">#让AI提前停止推理反而更优#</span></a><br><br>大模型推理越来越强，但也可能“想太多”，导致效率低、结论错。<br><br>华为联合中科院信工，选用不同规模的DeepSeek-R1-Distill-Qwen模型（1.5B, 7B, 14B, 32B），魔改出新方法DEER，能让模型在信息足够时“及时收手”。<br><br>DEER的关键是识别思维链中最有价值的临界点（Pearl Reasoning），在这个点就让模型提前结束推理、直接给出答案。实验表明：这种提前退出不仅提高准确率，还能减少31%到43%的生成长度。<br><br>研究发现，很多问题并不需要完整的推理链。比如在MATH-500数据集里，仅用20%的推理步骤就能保持60.8%的准确率。而GPQA这种更复杂的数据集，也有35.1%的样本能在少于一半路径时答对。<br><br>传统的固定退出点策略不够灵活，因此研究者设计了动态退出机制 DEER。它通过三步来判断退出时机：<br><br>1. 监控推理中出现的思路转换点（如“wait”等词）；<br><br>2. 在这些点强制诱导生成一个“试验性答案”；<br><br>3. 如果答案的置信度高，就提前终止思考；否则继续推理。<br><br>为了提高效率，DEER还引入了“分支并行”机制，让诱导、评估和推理并行进行，减少额外时延。<br><br>实际测试显示，DEER在多个benchmark和不同规模模型上都有效，准确率提升1.7-5.7个百分点，生成长度减少三分之一以上。尤其在人类代码生成任务中，生成长度减半，准确率还涨了。 <a href="https://weibo.com/ttarticle/p/show?id=2309405165246531895614" data-hide=""><span class="url-icon"><img style="width: 1rem;height: 1rem" src="https://h5.sinaimg.cn/upload/2015/09/25/3/timeline_card_small_article_default.png" referrerpolicy="no-referrer"></span><span class="surl-text">DeepSeek精度效率双提升，华为＆信工所提出思维链“提前退出”机制</span></a><img style="" src="https://tvax1.sinaimg.cn/large/006Fd7o3gy1i1bxh8td2qj30dv07tmxt.jpg" referrerpolicy="no-referrer"><br><br>

## AI 摘要

华为联合中科院信工所提出DEER方法，通过动态识别思维链临界点让AI提前终止推理。该方法在DeepSeek-R1-Distill-Qwen模型上测试显示：准确率提升1.7-5.7%，生成长度减少31%-43%。DEER采用三步机制（监测转换点、生成试验答案、置信度评估）和分支并行技术，在MATH-500等数据集中，部分问题仅需20%推理步骤即可保持60.8%准确率。特别在代码生成任务中，实现长度减半且准确率提升。研究证明适度"截断"推理反而能优化模型表现。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-05-12T05:04:16Z
- **目录日期**: 2025-05-12
