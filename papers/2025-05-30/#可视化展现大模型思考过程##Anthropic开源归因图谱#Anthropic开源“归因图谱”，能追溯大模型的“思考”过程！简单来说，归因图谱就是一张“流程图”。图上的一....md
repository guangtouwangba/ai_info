# #可视化展现大模型思考过程##Anthropic开源归因图谱#Anthropic开源“归因图谱”，能追溯大模型的“思考”过程！简单来说，归因图谱就是一张“流程图”。图上的一...

**URL**: https://weibo.com/6105753431/Puc93sYgL

## 原始摘要

<a href="https://m.weibo.cn/search?containerid=231522type%3D1%26t%3D10%26q%3D%23%E5%8F%AF%E8%A7%86%E5%8C%96%E5%B1%95%E7%8E%B0%E5%A4%A7%E6%A8%A1%E5%9E%8B%E6%80%9D%E8%80%83%E8%BF%87%E7%A8%8B%23&amp;extparam=%23%E5%8F%AF%E8%A7%86%E5%8C%96%E5%B1%95%E7%8E%B0%E5%A4%A7%E6%A8%A1%E5%9E%8B%E6%80%9D%E8%80%83%E8%BF%87%E7%A8%8B%23" data-hide=""><span class="surl-text">#可视化展现大模型思考过程#</span></a><a href="https://m.weibo.cn/search?containerid=231522type%3D1%26t%3D10%26q%3D%23Anthropic%E5%BC%80%E6%BA%90%E5%BD%92%E5%9B%A0%E5%9B%BE%E8%B0%B1%23&amp;extparam=%23Anthropic%E5%BC%80%E6%BA%90%E5%BD%92%E5%9B%A0%E5%9B%BE%E8%B0%B1%23" data-hide=""><span class="surl-text">#Anthropic开源归因图谱#</span></a><br><br>Anthropic开源“归因图谱”，能追溯大模型的“思考”过程！<br><br>简单来说，归因图谱就是一张“流程图”。图上的一个个小点（节点）代表了模型处理信息时识别出的各种“特征”，而连接这些点的线（边）则表示这些特征之间是怎么相互影响的。<br><br>它能图形化地展示模型为了给出一个特定的答案，都走了哪些计算步骤。<br><br>来看两个官方示例，咱们一看就明白：<br><br>1. 两跳推理<br>当你问Gemma 2（2B 模型）“达拉斯所在州的首府是什么”时，模型需要先想一下：达拉斯在德克萨斯州，然后再推导出德克萨斯州的首府是奥斯汀。<br><br>归因图谱就能把这种“先想一步”的过程清晰地画出来。【图2】<br><br>2. 多语言回路<br>另一个例子是Haiku模型，当你用不同语言问它“小的反义词是什么”时，它内部处理问题的“回路”竟然非常相似。<br><br>这可能说明 Haiku 在理解问题时，不是死板地看语言本身，而是先转化成了一种“不分语言”的通用概念来识别。（英文：【图3】中文：【图4】法语：【图5】）<br><br>Anthropic 这次开源的工具，普通人也能上手试试看！<br><br>想简单体验一下的同学，可以直接去Neuronpedia网站。选择一个模型和预设问题或提出新问题，它就能帮你生成对应的归因图谱，交互式地探索这些图谱的奥秘。【图6】<br><br>如果你是研究人员或者对代码感兴趣，可以直接去代码仓库深入研究。总的来说，这次开源能帮助大家：<br><br>- 在你支持的模型上，生成自己的归因图谱，追踪模型的“思考回路”。<br>- 在交互界面上，更直观地查看、标记和分享这些图谱。<br>- 通过调整图谱中的特征值，看看模型输出会有啥变化，从而验证自己的想法。<br><br>来试试看吧！<br>Neuronpedia：<a href="https://weibo.cn/sinaurl?u=https%3A%2F%2Fwww.neuronpedia.org%2Fgemma-2-2b%2Fgraph" data-hide=""><span class="url-icon"><img style="width: 1rem;height: 1rem" src="https://h5.sinaimg.cn/upload/2015/09/25/3/timeline_card_small_web_default.png" referrerpolicy="no-referrer"></span><span class="surl-text">网页链接</span></a><br>代码仓库：<a href="https://weibo.cn/sinaurl?u=https%3A%2F%2Fgithub.com%2Fsafety-research%2Fcircuit-tracer" data-hide=""><span class="url-icon"><img style="width: 1rem;height: 1rem" src="https://h5.sinaimg.cn/upload/2015/09/25/3/timeline_card_small_web_default.png" referrerpolicy="no-referrer"></span><span class="surl-text">网页链接</span></a><img style="" src="https://tvax1.sinaimg.cn/large/006Fd7o3gy1i1xd3ing90j31dc16un4u.jpg" referrerpolicy="no-referrer"><br><br><img style="" src="https://tvax2.sinaimg.cn/large/006Fd7o3gy1i1xd3mlkikj30zk0ho0y4.jpg" referrerpolicy="no-referrer"><br><br><img style="" src="https://tvax1.sinaimg.cn/large/006Fd7o3gy1i1xd4253mhj31o60qydp5.jpg" referrerpolicy="no-referrer"><br><br><img style="" src="https://tvax3.sinaimg.cn/large/006Fd7o3gy1i1xd448sm6j31d00ms79k.jpg" referrerpolicy="no-referrer"><br><br><img style="" src="https://tvax2.sinaimg.cn/large/006Fd7o3gy1i1xd45rxtbj31d40n0afp.jpg" referrerpolicy="no-referrer"><br><br><img style="" src="https://tvax1.sinaimg.cn/large/006Fd7o3gy1i1xd4a9qrej33gc1qw1ky.jpg" referrerpolicy="no-referrer"><br><br>

## AI 摘要

Anthropic开源了"归因图谱"工具，可可视化大语言模型的推理过程。该工具通过节点（特征）和边（关联）构成的流程图，展示模型回答问题的计算路径。示例显示：Gemma 2模型通过两跳推理回答"达拉斯所在州的首府"；Haiku模型处理多语言反义词问题时使用相似的内部回路，表明其可能采用语言无关的概念处理。该工具支持在Neuronpedia网站交互式探索，或通过GitHub代码库深入研究，允许用户追踪模型思维路径、标记图谱并验证特征影响。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-05-30T07:03:34Z
- **目录日期**: 2025-05-30
