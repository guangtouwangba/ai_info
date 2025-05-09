# #通义强化学习框架突破搜索引擎依赖##LLM搜索训练告别搜索引擎#利用强化学习提升LLMs的搜索能力，一定需要与真实搜索引擎交互吗？阿里巴巴通义实验室最新论文表...

**URL**: https://weibo.com/6105753431/PqZYRd3BM

## 原始摘要

<a href="https://m.weibo.cn/search?containerid=231522type%3D1%26t%3D10%26q%3D%23%E9%80%9A%E4%B9%89%E5%BC%BA%E5%8C%96%E5%AD%A6%E4%B9%A0%E6%A1%86%E6%9E%B6%E7%AA%81%E7%A0%B4%E6%90%9C%E7%B4%A2%E5%BC%95%E6%93%8E%E4%BE%9D%E8%B5%96%23&amp;extparam=%23%E9%80%9A%E4%B9%89%E5%BC%BA%E5%8C%96%E5%AD%A6%E4%B9%A0%E6%A1%86%E6%9E%B6%E7%AA%81%E7%A0%B4%E6%90%9C%E7%B4%A2%E5%BC%95%E6%93%8E%E4%BE%9D%E8%B5%96%23" data-hide=""><span class="surl-text">#通义强化学习框架突破搜索引擎依赖#</span></a><a href="https://m.weibo.cn/search?containerid=231522type%3D1%26t%3D10%26q%3D%23LLM%E6%90%9C%E7%B4%A2%E8%AE%AD%E7%BB%83%E5%91%8A%E5%88%AB%E6%90%9C%E7%B4%A2%E5%BC%95%E6%93%8E%23&amp;extparam=%23LLM%E6%90%9C%E7%B4%A2%E8%AE%AD%E7%BB%83%E5%91%8A%E5%88%AB%E6%90%9C%E7%B4%A2%E5%BC%95%E6%93%8E%23" data-hide=""><span class="surl-text">#LLM搜索训练告别搜索引擎#</span></a><br><br>利用强化学习提升LLMs的搜索能力，一定需要与真实搜索引擎交互吗？<br><br>阿里巴巴通义实验室最新论文表示，未必如此！【图1】<br><br>通常来说，采用真实搜索引擎进行训练的方式，会面临两个问题：<br><br>1. 文档质量不可控：搜索引擎返回的文档质量参差不齐，给训练过程带来噪声和不稳定性，影响模型训练效果。<br><br>2. API成本过高：强化学习训练需要频繁进行API调用，这会带来巨额费用，制约模型可扩展性。<br><br>为了应对这些挑战，研究团队创新性地提出了强化学习框架ZeroSearch，将LLMs转化为检索模块，无需与真实搜索引擎交互即可进行训练。【图2】<br><br>他们是如何做到的呢？关键在于以下四点：<br><br>1、轻量级监督微调（SFT）<br>- 收集LLMs与真实搜索引擎交互的轨迹，将产生正确答案的轨迹标记为正样本，产生错误答案的轨迹标记为负样本。<br>- 提取查询-文档对，并进行轻量级SFT，使其能够根据提示生成有用或噪声文档。<br>2、渐进式课程学习策略<br>- 在训练过程中，逐步增加生成文档的噪声比例，使模型逐渐适应更具挑战性的检索场景。<br>- 使用概率函数控制生成噪声文档的可能性，随着训练的进行，逐渐增加噪声文档的比例。<br>3、奖励设计：采用基于F1分数的奖励函数，专注于答案的准确性。<br>4、强化学习算法<br>- ZEROSEARCH与多种强化学习算法兼容，如近端策略优化（PPO）、组相对策略优化（GRPO）和Reinforce++等。<br>- 这些算法通过奖励信号指导模型的学习，使模型能够更好地掌握搜索策略。<br><br>从测试结果看，ZeroSearch的表现相当亮眼：<br><br>- 在多个问答基准数据集上均优于基线方法。【图3】<br>- 性能上超越了依赖真实搜索引擎的方法：【图3】<br>  - 70亿参数的检索模块已达到与真实搜索引擎相当的性能<br>  - 而140亿参数模块甚至实现了超越<br>- 在不同规模的基础模型和指令微调模型上均表现出良好的泛化性，并能兼容多种强化学习算法。【图4】<br><br>更多技术细节，欢迎点击论文链接查看➡️：<a href="https://weibo.cn/sinaurl?u=https%3A%2F%2Farxiv.org%2Fpdf%2F2505.04588" data-hide=""><span class="url-icon"><img style="width: 1rem;height: 1rem" src="https://h5.sinaimg.cn/upload/2015/09/25/3/timeline_card_small_web_default.png" referrerpolicy="no-referrer"></span><span class="surl-text">网页链接</span></a><br>项目主页：<a href="https://weibo.cn/sinaurl?u=https%3A%2F%2Falibaba-nlp.github.io%2FZeroSearch%2F" data-hide=""><span class="url-icon"><img style="width: 1rem;height: 1rem" src="https://h5.sinaimg.cn/upload/2015/09/25/3/timeline_card_small_web_default.png" referrerpolicy="no-referrer"></span><span class="surl-text">网页链接</span></a><img style="" src="https://tvax4.sinaimg.cn/large/006Fd7o3gy1i1929fspx4j30uc0wgqfo.jpg" referrerpolicy="no-referrer"><br><br><img style="" src="https://tvax2.sinaimg.cn/large/006Fd7o3gy1i1929h2g5oj30q80dq0wn.jpg" referrerpolicy="no-referrer"><br><br><img style="" src="https://tvax3.sinaimg.cn/large/006Fd7o3gy1i1929lq4r1j30qm0rg7fp.jpg" referrerpolicy="no-referrer"><br><br><img style="" src="https://tvax2.sinaimg.cn/large/006Fd7o3gy1i1929kkbt6j30qm07iq7g.jpg" referrerpolicy="no-referrer"><br><br>

## AI 摘要

阿里巴巴通义实验室提出新型强化学习框架ZeroSearch，突破传统LLM搜索训练必须依赖真实搜索引擎的限制。该框架通过轻量级监督微调、渐进式课程学习策略和基于F1的奖励设计，使LLM自身成为检索模块。实验显示，70亿参数模型性能媲美真实搜索引擎，140亿参数模型甚至实现超越，同时解决了传统方法面临的文档质量不可控和API成本过高问题。该技术在不同规模模型上均展现良好泛化性，兼容多种强化学习算法，为LLM搜索训练提供了更高效经济的解决方案。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-05-09T08:04:03Z
- **目录日期**: 2025-05-09
