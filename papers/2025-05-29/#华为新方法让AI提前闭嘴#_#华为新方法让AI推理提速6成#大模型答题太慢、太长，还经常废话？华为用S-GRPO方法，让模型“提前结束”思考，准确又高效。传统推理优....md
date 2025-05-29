# #华为新方法让AI提前闭嘴# #华为新方法让AI推理提速6成#大模型答题太慢、太长，还经常废话？华为用S-GRPO方法，让模型“提前结束”思考，准确又高效。传统推理优...

**URL**: https://weibo.com/6105753431/Pu50WEIkm

## 原始摘要

<a href="https://m.weibo.cn/search?containerid=231522type%3D1%26t%3D10%26q%3D%23%E5%8D%8E%E4%B8%BA%E6%96%B0%E6%96%B9%E6%B3%95%E8%AE%A9AI%E6%8F%90%E5%89%8D%E9%97%AD%E5%98%B4%23&amp;extparam=%23%E5%8D%8E%E4%B8%BA%E6%96%B0%E6%96%B9%E6%B3%95%E8%AE%A9AI%E6%8F%90%E5%89%8D%E9%97%AD%E5%98%B4%23" data-hide=""><span class="surl-text">#华为新方法让AI提前闭嘴#</span></a> <a href="https://m.weibo.cn/search?containerid=231522type%3D1%26t%3D10%26q%3D%23%E5%8D%8E%E4%B8%BA%E6%96%B0%E6%96%B9%E6%B3%95%E8%AE%A9AI%E6%8E%A8%E7%90%86%E6%8F%90%E9%80%9F6%E6%88%90%23&amp;extparam=%23%E5%8D%8E%E4%B8%BA%E6%96%B0%E6%96%B9%E6%B3%95%E8%AE%A9AI%E6%8E%A8%E7%90%86%E6%8F%90%E9%80%9F6%E6%88%90%23" data-hide=""><span class="surl-text">#华为新方法让AI推理提速6成#</span></a><br><br>大模型答题太慢、太长，还经常废话？华为用S-GRPO方法，让模型“提前结束”思考，准确又高效。<br><br>传统推理优化GRPO喜欢“拖到底”，S-GRPO反其道而行：让模型在思考过程中可以中途“打住”，一旦答对，越早停奖励越高。<br><br>S-GRPO可看作是“串行分组 + 衰减奖励”，总体分三步走：<br><br>1. 先完整思考：模型像平时那样推完整条路径。<br><br>2. 再切片思考：中间随时打断，训练模型在不同深度都能答题。<br><br>3. 按早答奖励：答对就看退出早不早，越早越奖，答错就一毛不给。<br><br>这套方法训练出的模型更聪明，能提前结束思考，节省60%计算量，答案还更准确。<br><br>实测对比显示，在数学和科学推理五大任务上，S-GRPO的表现大幅领先传统方法。不仅推理路径更短，答对率也更高。而且在长度受限的情况下优势更明显，几乎用一半“预算”就能达到同等甚至更好的效果。详细内容请点击文章： <a href="https://weibo.com/ttarticle/p/show?id=2309405171691193237578" data-hide=""><span class="url-icon"><img style="width: 1rem;height: 1rem" src="https://h5.sinaimg.cn/upload/2015/09/25/3/timeline_card_small_article_default.png" referrerpolicy="no-referrer"></span><span class="surl-text">华为攻克AI推理「想太多」问题！新方法让大模型推理提速60%。</span></a><img style="" src="https://tvax4.sinaimg.cn/large/006Fd7o3gy1i1whmieijij30m00cejsj.jpg" referrerpolicy="no-referrer"><br><br>

## AI 摘要

华为提出S-GRPO新方法，解决大模型推理"想太多"问题。该方法通过"串行分组+衰减奖励"机制，允许模型在推理过程中提前终止：完整思考后训练模型在不同深度答题，越早答对奖励越高。实验显示，S-GRPO在数学和科学推理任务中，能节省60%计算量，同时提高准确率。与传统方法相比，其推理路径更短，在有限计算资源下表现更优，用一半预算即可达到更好效果。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-05-29T13:12:11Z
- **目录日期**: 2025-05-29
