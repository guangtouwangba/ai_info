# #大模型会跳跃思考了##软推理让大模型像人一样思考#大模型“一个字一个字”地推理的日子，或许要结束了。SimularAI与微软DeepSpeed联合提出的Soft Thinking方法...

**URL**: https://weibo.com/6105753431/Ptid3ricq

## 原始摘要

<a href="https://m.weibo.cn/search?containerid=231522type%3D1%26t%3D10%26q%3D%23%E5%A4%A7%E6%A8%A1%E5%9E%8B%E4%BC%9A%E8%B7%B3%E8%B7%83%E6%80%9D%E8%80%83%E4%BA%86%23&amp;extparam=%23%E5%A4%A7%E6%A8%A1%E5%9E%8B%E4%BC%9A%E8%B7%B3%E8%B7%83%E6%80%9D%E8%80%83%E4%BA%86%23" data-hide=""><span class="surl-text">#大模型会跳跃思考了#</span></a><a href="https://m.weibo.cn/search?containerid=231522type%3D1%26t%3D10%26q%3D%23%E8%BD%AF%E6%8E%A8%E7%90%86%E8%AE%A9%E5%A4%A7%E6%A8%A1%E5%9E%8B%E5%83%8F%E4%BA%BA%E4%B8%80%E6%A0%B7%E6%80%9D%E8%80%83%23&amp;extparam=%23%E8%BD%AF%E6%8E%A8%E7%90%86%E8%AE%A9%E5%A4%A7%E6%A8%A1%E5%9E%8B%E5%83%8F%E4%BA%BA%E4%B8%80%E6%A0%B7%E6%80%9D%E8%80%83%23" data-hide=""><span class="surl-text">#软推理让大模型像人一样思考#</span></a><br><br>大模型“一个字一个字”地推理的日子，或许要结束了。<br><br>SimularAI与微软DeepSpeed联合提出的Soft Thinking方法，通过“软推理”突破传统思维链（CoT）的限制，让AI像人类一样进行抽象思考。<br><br>传统CoT依赖离散token推理，路径指数级增长，效率低且不够灵活。Soft Thinking创新性地引入“连续概念空间”和“概念token”，用概率分布代替单一token选择，实现多条推理路径的隐式聚合。不仅更抽象灵活，还减少22.4% token使用量。<br><br>举个例子，遇到“43×34”的题目，Soft Thinking不再选“直接乘”或“分解后乘”之一，而是能同时保留两种思路。通过词向量加权，模型能在多种语义间平滑过渡，像人脑一样跳跃思考。<br><br>为防模型“卡壳”，研究者还设计了Cold Stop机制，监测模型“自信程度”，自如判断何时停止推理。<br><br>测试结果显示，QwQ-32B模型准确率最多提升2.48%，在AIME 2024数据集上更提升6.45%。相比其他无训练或平均嵌入策略，Soft Thinking在效率与准确率上表现更优。<br><br>这种即插即用的“软推理”，或将成为未来大模型推理的标配方式。感兴趣的小伙伴点击：<a href="https://weibo.cn/sinaurl?u=https%3A%2F%2Fmp.weixin.qq.com%2Fs%2Fa12P9YSNcII565BA7NBB1Q" data-hide=""><span class="url-icon"><img style="width: 1rem;height: 1rem" src="https://h5.sinaimg.cn/upload/2015/09/25/3/timeline_card_small_web_default.png" referrerpolicy="no-referrer"></span><span class="surl-text">打破思维链推理瓶颈！“软推理”让大模型学会人类抽象能力，token使用量还更少了</span></a><img style="" src="https://tvax3.sinaimg.cn/large/006Fd7o3ly1i1qi3sy0t1j30zk0rodq6.jpg" referrerpolicy="no-referrer"><br><br><img style="" src="https://tvax4.sinaimg.cn/large/006Fd7o3ly1i1qi3ui4cmj30p00go0xd.jpg" referrerpolicy="no-referrer"><br><br><img style="" src="https://tvax2.sinaimg.cn/large/006Fd7o3ly1i1qi3wrpmxj30p00lewk6.jpg" referrerpolicy="no-referrer"><br><br><img style="" src="https://tvax2.sinaimg.cn/large/006Fd7o3ly1i1qi3zugxuj30zk0senbd.jpg" referrerpolicy="no-referrer"><br><br>

## AI 摘要

SimularAI与微软DeepSpeed联合提出的"软推理"(Soft Thinking)方法突破传统思维链(CoT)限制，通过引入连续概念空间和概念token，用概率分布代替单一token选择，实现多推理路径隐式聚合。该方法使大模型能像人类一样进行抽象跳跃思考，在保持43×34这类问题多种解法思路的同时，减少22.4%的token使用量。测试显示QwQ-32B模型准确率最高提升2.48%，在AIME 2024数据集上提升6.45%。创新的Cold Stop机制还能动态判断推理终止时机，使AI推理更接近人类思维方式。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-05-25T14:04:32Z
- **目录日期**: 2025-05-25
