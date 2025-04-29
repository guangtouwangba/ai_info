# #清华等团队让RAG性能提升20%##清华等团队发布DeepNote深度笔记#清华、中科大等团队，发布了DeepNote深度笔记，让RAG性能暴增20%。这是一种全新的适应式RAG方法...

**URL**: https://weibo.com/6105753431/PpuSeEvng

## 原始摘要

<a href="https://m.weibo.cn/search?containerid=231522type%3D1%26t%3D10%26q%3D%23%E6%B8%85%E5%8D%8E%E7%AD%89%E5%9B%A2%E9%98%9F%E8%AE%A9RAG%E6%80%A7%E8%83%BD%E6%8F%90%E5%8D%8720%25%23&amp;extparam=%23%E6%B8%85%E5%8D%8E%E7%AD%89%E5%9B%A2%E9%98%9F%E8%AE%A9RAG%E6%80%A7%E8%83%BD%E6%8F%90%E5%8D%8720%25%23" data-hide=""><span class="surl-text">#清华等团队让RAG性能提升20%#</span></a><a href="https://m.weibo.cn/search?containerid=231522type%3D1%26t%3D10%26q%3D%23%E6%B8%85%E5%8D%8E%E7%AD%89%E5%9B%A2%E9%98%9F%E5%8F%91%E5%B8%83DeepNote%E6%B7%B1%E5%BA%A6%E7%AC%94%E8%AE%B0%23&amp;extparam=%23%E6%B8%85%E5%8D%8E%E7%AD%89%E5%9B%A2%E9%98%9F%E5%8F%91%E5%B8%83DeepNote%E6%B7%B1%E5%BA%A6%E7%AC%94%E8%AE%B0%23" data-hide=""><span class="surl-text">#清华等团队发布DeepNote深度笔记#</span></a><br><br>清华、中科大等团队，发布了DeepNote深度笔记，让RAG性能暴增20%。<br><br>这是一种全新的适应式RAG方法，首次引入“笔记”作为知识载体，使知识探索更深入。<br><br>在所有任务上，DeepNote比主流RAG方法性能提升高达20.1%，即使在中小参数模型下也能强力泛化。<br><br>传统RAG方法只支持一次性检索，面对复杂推理任务力有未逮。<br><br>多轮检索虽能缓解问题，但会引入大量无关噪声，降低回答质量。<br><br>自适应RAG虽然引入动态决策，但仍面临检索生成耦合过紧、检索策略决策不足等痛点，缺乏真正的信息生长能力。<br><br>而DeepNote围绕“笔记”展开知识积累，核心包含三个阶段：<br><br>- 笔记初始化：初始问题和首次检索内容生成笔记，作为后续依据。<br>- 基于笔记的适应式检索：判断新信息是否有价值，有就更新笔记，否则终止检索。<br>- 基于最佳笔记生成答案：用积累的知识统一作答，保证逻辑连贯、来源清晰。<br><br>这种设计，能模拟人类查阅笔记、逐步解决问题的过程。<br><br>而且该方法是目前唯一同时在自适应检索控制、知识积累更新、模型优化三大维度实现突破的方法，彻底打破传统“一问一答”的局限。<br><br>实验结果方面，在HotpotQA、2WikiMQA、MusiQue、ASQA、StrategyQA等五大数据集上，DeepNote在所有任务中全面超越现有方法，整体性能提升显著。<br><br>同时，团队构建了DNAlign数据集，并结合DPO优化技术，进一步提升模型指令遵循能力。详情请见文章： <a href="https://weibo.com/ttarticle/p/show?id=2309405160768827162768" data-hide=""><span class="url-icon"><img style="width: 1rem;height: 1rem" src="https://h5.sinaimg.cn/upload/2015/09/25/3/timeline_card_small_article_default.png" referrerpolicy="no-referrer"></span><span class="surl-text">RAG性能暴增20%！清华等推出“以笔记为中心”的深度检索增强生成框架</span></a><img style="" src="https://tvax4.sinaimg.cn/large/006Fd7o3gy1i0xi4m4qwsj30ka0bface.jpg" referrerpolicy="no-referrer"><br><br>

## AI 摘要

清华大学等团队提出新型RAG框架DeepNote，通过引入"笔记"作为动态知识载体，实现检索性能提升20.1%。该方法创新性地采用三阶段流程：笔记初始化、适应式检索更新、基于最优笔记生成答案，模拟人类渐进式知识积累过程。实验在五大QA数据集上全面超越现有方法，尤其在复杂推理任务中表现突出。DeepNote首次同时实现自适应检索控制、知识动态更新和模型优化三大突破，并配套发布DNAlign数据集。该技术显著改善了传统RAG一次性检索的局限性，为知识密集型任务提供新范式。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-04-29T08:02:58Z
- **目录日期**: 2025-04-29
