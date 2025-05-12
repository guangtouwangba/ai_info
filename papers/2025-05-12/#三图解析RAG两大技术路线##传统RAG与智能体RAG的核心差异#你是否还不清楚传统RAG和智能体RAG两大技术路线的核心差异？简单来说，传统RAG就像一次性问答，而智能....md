# #三图解析RAG两大技术路线##传统RAG与智能体RAG的核心差异#你是否还不清楚传统RAG和智能体RAG两大技术路线的核心差异？简单来说，传统RAG就像一次性问答，而智能...

**URL**: https://weibo.com/6105753431/PrsLDnC28

## 原始摘要

<a href="https://m.weibo.cn/search?containerid=231522type%3D1%26t%3D10%26q%3D%23%E4%B8%89%E5%9B%BE%E8%A7%A3%E6%9E%90RAG%E4%B8%A4%E5%A4%A7%E6%8A%80%E6%9C%AF%E8%B7%AF%E7%BA%BF%23&amp;extparam=%23%E4%B8%89%E5%9B%BE%E8%A7%A3%E6%9E%90RAG%E4%B8%A4%E5%A4%A7%E6%8A%80%E6%9C%AF%E8%B7%AF%E7%BA%BF%23" data-hide=""><span class="surl-text">#三图解析RAG两大技术路线#</span></a><a href="https://m.weibo.cn/search?containerid=231522type%3D1%26t%3D10%26q%3D%23%E4%BC%A0%E7%BB%9FRAG%E4%B8%8E%E6%99%BA%E8%83%BD%E4%BD%93RAG%E7%9A%84%E6%A0%B8%E5%BF%83%E5%B7%AE%E5%BC%82%23&amp;extparam=%23%E4%BC%A0%E7%BB%9FRAG%E4%B8%8E%E6%99%BA%E8%83%BD%E4%BD%93RAG%E7%9A%84%E6%A0%B8%E5%BF%83%E5%B7%AE%E5%BC%82%23" data-hide=""><span class="surl-text">#传统RAG与智能体RAG的核心差异#</span></a><br><br>你是否还不清楚传统RAG和智能体RAG两大技术路线的核心差异？<br><br>简单来说，传统RAG就像一次性问答，而智能体RAG更像是会反复思考的对话。下面用三张图带你看懂这个技术升级。<br><br>在传统RAG的检索过程中，用户提出的问题会被转换为向量，根据向量相似度，从数据库中查询最相关的文本片段，生成模型会结合检索到的上下文和原始查询生成答案。【图1】<br><br>但这样做有三个明显短板：<br>1. 单次检索单次生成：如果第一次找的资料不够，系统不会自己补充<br>2. 缺乏复杂推理能力：遇到需要多步推理的问题就容易卡壳<br>3. 处理方式固定：不管什么问题都用同一种方法解决<br><br>而智能体RAG就通过引入智能体，来解决这三个短板：【图2】<br><br>1. 步骤1-2：先帮你优化问题，比如自动改正错别字<br>2. 步骤3-8：判断是否需要补充上下文。若不需要，将优化后的查询直接发送给大模型；若需要，智能体将会自己去找合适的参考资料<br>3. 步骤9-12：获得初始响应后，智能体验证答案相关性：若相关，返回最终答案；若不相关，返回步骤1重新执行<br><br>这个“思考-验证-调整”的过程会循环几次，直到给出满意答案或确认真的回答不了。<br><br>这就相当于给系统装了个“质检员”，RAG的鲁棒性得到了提升，每个环节都朝着给出好答案的目标努力。<img style="" src="https://tvax4.sinaimg.cn/large/006Fd7o3gy1i1clck8zh4g30zk0f6k3u.gif" referrerpolicy="no-referrer"><br><br><img style="" src="https://tvax1.sinaimg.cn/large/006Fd7o3gy1i1clct7zbrg30zk0oiarp.gif" referrerpolicy="no-referrer"><br><br><img style="" src="https://tvax4.sinaimg.cn/large/006Fd7o3gy1i1clcl0536g30ue0zknmg.gif" referrerpolicy="no-referrer"><br><br>

## AI 摘要

传统RAG通过单次检索生成答案，存在检索不足、推理能力弱和处理方式单一等局限。智能体RAG引入智能体机制，实现多轮优化：先修正查询问题，动态判断是否需要补充检索，验证答案相关性并循环调整，直至输出满意结果。这种"思考-验证-调整"的闭环流程显著提升了系统的鲁棒性和复杂问题处理能力，使RAG从一次性问答升级为具备自我修正能力的智能对话系统。三张示意图分别展示了传统流程、智能体决策路径和整体技术升级对比。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-05-12T08:03:35Z
- **目录日期**: 2025-05-12
