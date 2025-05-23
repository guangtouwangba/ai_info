# #超长推理还能节省计算##提升模型推理效率新方法#大模型推理效率怎么提升？Salesforce AI Research开源了两个方案：Elastic Reasoning和Fractured Sampling。Ela...

**URL**: https://weibo.com/6105753431/Pta2G6CiD

## 原始摘要

<a href="https://m.weibo.cn/search?containerid=231522type%3D1%26t%3D10%26q%3D%23%E8%B6%85%E9%95%BF%E6%8E%A8%E7%90%86%E8%BF%98%E8%83%BD%E8%8A%82%E7%9C%81%E8%AE%A1%E7%AE%97%23&amp;extparam=%23%E8%B6%85%E9%95%BF%E6%8E%A8%E7%90%86%E8%BF%98%E8%83%BD%E8%8A%82%E7%9C%81%E8%AE%A1%E7%AE%97%23" data-hide=""><span class="surl-text">#超长推理还能节省计算#</span></a><a href="https://m.weibo.cn/search?containerid=231522type%3D1%26t%3D10%26q%3D%23%E6%8F%90%E5%8D%87%E6%A8%A1%E5%9E%8B%E6%8E%A8%E7%90%86%E6%95%88%E7%8E%87%E6%96%B0%E6%96%B9%E6%B3%95%23&amp;extparam=%23%E6%8F%90%E5%8D%87%E6%A8%A1%E5%9E%8B%E6%8E%A8%E7%90%86%E6%95%88%E7%8E%87%E6%96%B0%E6%96%B9%E6%B3%95%23" data-hide=""><span class="surl-text">#提升模型推理效率新方法#</span></a><br><br>大模型推理效率怎么提升？Salesforce AI Research开源了两个方案：Elastic Reasoning和Fractured Sampling。<br><br>Elastic Reasoning主打“想多少、答多少”。它不再让模型无限生成推理链，而是把“思考”和“解题”分开设定token预算，用完思考预算后立刻转入答题。这种方式不仅让答案更完整，效率也提升了30%。比如E1-Math-1.5B模型，就在准确率领先的同时，用更少训练资源达到了更优表现。<br><br>Fractured Sampling则鼓励模型“少想早答”。它把推理过程在时间上切割，对每条推理链做“提前停想”，沿三个维度采样：思考路径数n、每路径答案数m、思考深度H。实验发现，提升H带来的性能收益远高于其他两个维度。DeepSeek-R1等模型在应用后，token减少约20%，精度还能升。<br><br>两个方法都有一个目标：在推理预算有限时，也能交出漂亮答卷。不仅让大模型“又快又准”，还在编程和数学任务中表现出色。 <a href="https://weibo.com/ttarticle/p/show?id=2309405169501179347409" data-hide=""><span class="url-icon"><img style="width: 1rem;height: 1rem" src="https://h5.sinaimg.cn/upload/2015/09/25/3/timeline_card_small_article_default.png" referrerpolicy="no-referrer"></span><span class="surl-text">超长推理还能节省计算！Salesforce教LLM边想边省，显著提升数学编程</span></a><img style="" src="https://tvax3.sinaimg.cn/large/006Fd7o3gy1i1pi2r98lmj30n40d0jsd.jpg" referrerpolicy="no-referrer"><br><br>

## AI 摘要

Salesforce AI Research 开源了两种提升大模型推理效率的方法：Elastic Reasoning 和 Fractured Sampling。Elastic Reasoning 将“思考”和“答题”分开设定 token 预算，用完思考预算后立即答题，效率提升 30%，同时保持高准确率。Fractured Sampling 则切割推理过程，通过提前终止部分推理路径（重点优化思考深度 H），减少 20% token 使用且精度提升。两种方法均优化了计算资源分配，使模型在数学和编程任务中实现高效精准的推理。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-05-23T10:04:16Z
- **目录日期**: 2025-05-23
