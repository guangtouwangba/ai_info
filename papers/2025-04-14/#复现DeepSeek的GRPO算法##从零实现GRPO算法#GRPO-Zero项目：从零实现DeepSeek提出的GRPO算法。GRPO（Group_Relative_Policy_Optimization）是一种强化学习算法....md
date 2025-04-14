# #复现DeepSeek的GRPO算法##从零实现GRPO算法#GRPO-Zero项目：从零实现DeepSeek提出的GRPO算法。GRPO（Group Relative Policy Optimization）是一种强化学习算法...

**URL**: https://weibo.com/6105753431/PnbLIkN51

## 原始摘要

<a href="https://m.weibo.cn/search?containerid=231522type%3D1%26t%3D10%26q%3D%23%E5%A4%8D%E7%8E%B0DeepSeek%E7%9A%84GRPO%E7%AE%97%E6%B3%95%23&amp;extparam=%23%E5%A4%8D%E7%8E%B0DeepSeek%E7%9A%84GRPO%E7%AE%97%E6%B3%95%23" data-hide=""><span class="surl-text">#复现DeepSeek的GRPO算法#</span></a><a href="https://m.weibo.cn/search?containerid=231522type%3D1%26t%3D10%26q%3D%23%E4%BB%8E%E9%9B%B6%E5%AE%9E%E7%8E%B0GRPO%E7%AE%97%E6%B3%95%23&amp;extparam=%23%E4%BB%8E%E9%9B%B6%E5%AE%9E%E7%8E%B0GRPO%E7%AE%97%E6%B3%95%23" data-hide=""><span class="surl-text">#从零实现GRPO算法#</span></a><br><br>GRPO-Zero项目：从零实现DeepSeek提出的GRPO算法。<br><br>GRPO（Group Relative Policy Optimization）是一种强化学习算法，区别于传统RLHF，它不依赖value模型，也不用参考模型算KL散度，它的训练逻辑是这样的：<br><br>- 每轮训练先从数据集中抽几道题；<br>- 模型自由发挥，给每题随机“构造”几个回答；<br>- 为每个回答打分，例如判断是否答对、格式是否正确；<br>- 在每题内部对得分进行标准化，计算每个回答的“优势”；<br>- 用这个“优势”来更新策略模型，使模型更倾向于生成高分回答；<br>- 重复上述过程。<br><br>而GRPO-Zero项目从零实现了GRPO算法，它结构简单，依赖轻量，仅需PyTorch即可运行，甚至不依赖Transformers库，非常适合小团队快速上手。<br><br>实际训练时，一张A40显卡（48G显存）即可胜任，也可以部署在云端运行。<br><br>项目默认的训练任务是“倒计时题目”，也就是给出一组数字和目标结果，要求模型生成一个表达式，算出目标值。例如：<br><br>“给出数字1 2 3 4，目标是 11”，  <br>正确输出可以是：1 + (2 * 3) + 4。<br><br>模型还得按特定格式输出才能得分，标准答案形如：<br><br>逐步推理过程  <br>最终答案  <br><br>奖励设计也很直观，用完所有数字且计算正确就遵循：<br><br>- 格式对 ➜ 奖励 0.1 分；<br>- 答案对 ➜ 奖励 1 分；<br><br>训练使用的是Qwen2.5-3B-Instruct模型，项目提供了运行的完整代码，一键即可启动训练。【图2】<br><br>如果你正在找一个跑得动、看得懂、还能自由改的RLHF（Reinforcement Learning from Human Feedback，人类反馈强化学习）项目，GRPO-Zero会是你的起点。<br><br>感兴趣的小伙伴可以点击项目地址：<a href="https://weibo.cn/sinaurl?u=https%3A%2F%2Fgithub.com%2Fpolicy-gradient%2FGRPO-Zero" data-hide=""><span class="url-icon"><img style="width: 1rem;height: 1rem" src="https://h5.sinaimg.cn/upload/2015/09/25/3/timeline_card_small_web_default.png" referrerpolicy="no-referrer"></span><span class="surl-text">网页链接</span></a><img style="" src="https://tvax1.sinaimg.cn/large/006Fd7o3gy1i0g3g3w56bj31w619ih2l.jpg" referrerpolicy="no-referrer"><br><br><img style="" src="https://tvax3.sinaimg.cn/large/006Fd7o3gy1i0g3g598uyj30yy0ha430.jpg" referrerpolicy="no-referrer"><br><br>

## AI 摘要

GRPO-Zero项目从零实现了DeepSeek提出的GRPO（Group Relative Policy Optimization）强化学习算法。该算法无需价值模型或KL散度计算，通过抽样问题、生成多个回答、评分并标准化优势来更新策略模型。项目结构简单，仅需PyTorch，支持A40显卡训练，默认任务为数字表达式求解（如用1 2 3 4算出11）。奖励机制结合格式（0.1分）和答案正确性（1分），基于Qwen2.5-3B-Instruct模型，提供完整代码，适合小团队快速入门RLHF。GitHub链接：https://github.com/policy-gradient/GRPO-Zero

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-04-14T06:02:00Z
- **目录日期**: 2025-04-14
